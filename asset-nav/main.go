//go:build wasip1

package main

import (
	"asset-nav/contracts/evm/src/generated/asset_nav_contracts"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/scheduler/cron"
	"github.com/smartcontractkit/cre-sdk-go/cre"
	"github.com/smartcontractkit/cre-sdk-go/cre/wasm"
)

type EvmConfig struct {
	ChainSelector   uint64 `json:"chainSelector"`
	StorageAddress  string `json:"storageAddress"`
	ConsumerAddress string `json:"consumerAddress"`
	GasLimit        uint64 `json:"gasLimit"`
}

type Config struct {
	Schedule string      `json:"schedule"`
	ApiUrl   string      `json:"apiUrl"`
	Evms     []EvmConfig `json:"evms"`
}

type MyResult struct {
	OffchainValue        string
	PreviousOnchainValue string
	UpdatedValue         string
	TxHash               string
}

func InitWorkflow(config *Config, logger *slog.Logger, secretsProvider cre.SecretsProvider) (cre.Workflow[*Config], error) {
	return cre.Workflow[*Config]{
		cre.Handler(cron.Trigger(&cron.Config{Schedule: config.Schedule}), onCronTrigger),
	}, nil
}

func onCronTrigger(config *Config, runtime cre.Runtime, trigger *cron.Payload) (*MyResult, error) {
	logger := runtime.Logger()
	evmConfig := config.Evms[0]

	// Step 1: Fetch offchain data
	client := &http.Client{}
	navPricePromise := http.SendRequest(config, runtime, client, fetchResult, cre.ConsensusMedianAggregation[*big.Int]())
	offchainValue, err := navPricePromise.Await()
	if err != nil {
		return nil, err
	}

	usdFloat := new(big.Float).SetInt(offchainValue)
	usdFloat.Quo(usdFloat, big.NewFloat(1e18))
	usd64, _ := usdFloat.Float64()
	offchainPrettyStr := fmt.Sprintf("%.2f", usd64)
	pretty := map[string]map[string]float64{"ethereum": {"usd": usd64}}
	prettyJSON, _ := json.Marshal(pretty)
	logger.Info("Successfully fetched offchain value", "json", string(prettyJSON))

	// Step 2: Read onchain data using the binding for the Storage contract
	evmClient := &evm.Client{ChainSelector: evmConfig.ChainSelector}

	storageAddress := common.HexToAddress(evmConfig.ConsumerAddress)
	storageContract, err := asset_nav_contracts.NewAssetNavContracts(evmClient, storageAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract instance: %w", err)
	}

	onchainValue, err := storageContract.AssetNAV(runtime, big.NewInt(-3)).Await()
	if err != nil {
		return nil, fmt.Errorf("failed to read onchain value: %w", err)
	}

	// Convert onchain value (scaled by 100) to human-readable format
	onchainFloat := new(big.Float).SetInt(onchainValue)
	onchainFloat.Quo(onchainFloat, big.NewFloat(100))
	onchainUsd, _ := onchainFloat.Float64()
	onchainPretty := map[string]map[string]float64{"ethereum": {"usd": onchainUsd}}
	onchainPrettyJSON, _ := json.Marshal(onchainPretty)
	logger.Info("Successfully read onchain value", "raw", onchainValue, "json", string(onchainPrettyJSON))

	// Step 3: Prepare scaled offchain result
	rounder := big.NewInt(5000000000000000) // 5e15
	tmp := new(big.Int).Add(offchainValue, rounder)
	denom := new(big.Int).SetUint64(10000000000000000) // 1e16
	finalResultScaled := new(big.Int).Quo(tmp, denom)

	// Convert finalResultScaled (scaled by 100) to human-readable format
	finalResultFloat := new(big.Float).SetInt(finalResultScaled)
	finalResultFloat.Quo(finalResultFloat, big.NewFloat(100))
	finalResultUsd, _ := finalResultFloat.Float64()
	finalResultPretty := fmt.Sprintf("%.2f", finalResultUsd)
	finalPretty := map[string]map[string]string{"ethereum": {"usd": finalResultPretty}}
	finalPrettyJSON, _ := json.Marshal(finalPretty)
	logger.Info("Final calculated result to submit onchain", "raw", finalResultScaled, "json", string(finalPrettyJSON))

	// Step 4: Generate and submit CRE report using new WriteReport pattern
	txHash, err := generateAndSendReport(runtime, evmClient, evmConfig, finalResultScaled)
	if err != nil {
		return nil, fmt.Errorf("failed to update result: %w", err)
	}

	// Step 5: Return and log result
	onchainPrettyStr := fmt.Sprintf("%.2f", onchainUsd)
	finalWorkflowResult := &MyResult{
		OffchainValue:        offchainPrettyStr,
		PreviousOnchainValue: onchainPrettyStr,
		UpdatedValue:         finalResultPretty,
		TxHash:               txHash,
	}

	logger.Info("Workflow finished successfully", "result", finalWorkflowResult)
	return finalWorkflowResult, nil
}

func fetchResult(config *Config, logger *slog.Logger, sendRequester *http.SendRequester) (*big.Int, error) {
	req := &http.Request{Url: config.ApiUrl, Method: "GET"}
	resp, err := sendRequester.SendRequest(req).Await()
	if err != nil {
		return nil, fmt.Errorf("failed to get API response: %w", err)
	}

	var result struct {
		Ethereum struct {
			USD float64 `json:"usd"`
		} `json:"ethereum"`
	}

	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	val := new(big.Float).SetFloat64(result.Ethereum.USD)
	val.Mul(val, big.NewFloat(1e18))
	bigIntVal, _ := val.Int(nil)

	return bigIntVal, nil
}

func generateAndSendReport(runtime cre.Runtime, evmClient *evm.Client, evmConfig EvmConfig, finalResult *big.Int) (string, error) {
	logger := runtime.Logger()

	// Step 1: Encode payload
	intType, err := abi.NewType("int256", "", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create ABI type: %w", err)
	}
	args := abi.Arguments{{Type: intType}}
	calldata, err := args.Pack(finalResult)
	if err != nil {
		return "", fmt.Errorf("failed to ABI-encode payload: %w", err)
	}
	logger.Info("Encoded payload", "hex", fmt.Sprintf("0x%x", calldata))

	// Step 2: Generate CRE report
	reportPromise := runtime.GenerateReport(&cre.ReportRequest{
		EncodedPayload: calldata,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	report, err := reportPromise.Await()
	if err != nil {
		return "", fmt.Errorf("failed to generate report: %w", err)
	}
	logger.Info("CRE report generated successfully")

	// Step 3: Send report to blockchain
	receiverAddress := common.HexToAddress(evmConfig.ConsumerAddress)
	gasConfig := &evm.GasConfig{GasLimit: evmConfig.GasLimit}

	writePromise := evmClient.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver:  receiverAddress.Bytes(),
		Report:    report,
		GasConfig: gasConfig,
	})

	logger.Info("Submitting CRE report onchain...")
	resp, err := writePromise.Await()
	if err != nil {
		return "", fmt.Errorf("failed to submit report: %w", err)
	}

	if resp.TxStatus != evm.TxStatus_TX_STATUS_SUCCESS {
		errorMsg := "unknown error"
		if resp.ErrorMessage != nil {
			errorMsg = *resp.ErrorMessage
		}
		return "", fmt.Errorf("transaction failed with status %v: %s", resp.TxStatus, errorMsg)
	}

	txHash := fmt.Sprintf("0x%x", resp.TxHash)
	logger.Info("Report submitted successfully", "txHash", txHash, "fee", resp.TransactionFee)
	return txHash, nil
}

func main() {
	wasm.NewRunner(cre.ParseJSON[Config]).Run(InitWorkflow)
}
