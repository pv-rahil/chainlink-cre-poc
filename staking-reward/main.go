//go:build wasip1

package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"

	"onchain-calculator/contracts/evm/src/generated/storage"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/scheduler/cron"
	"github.com/smartcontractkit/cre-sdk-go/cre"
	"github.com/smartcontractkit/cre-sdk-go/cre/wasm"
)

// The EvmConfig is updated from Part 3 with new fields for the write operation.
type EvmConfig struct {
	ChainSelector             uint64 `json:"chainSelector"`
	StorageAddress            string `json:"storageAddress"`
	CalculatorConsumerAddress string `json:"calculatorConsumerAddress"`
	GasLimit                  uint64 `json:"gasLimit"`
}

type Config struct {
	Schedule string      `json:"schedule"`
	ApiUrl   string      `json:"apiUrl"`
	Evms     []EvmConfig `json:"evms"`
}

// MyResult struct now holds all the outputs of our workflow.
type MyResult struct {
	OffchainValue *big.Int
	OnchainValue  *big.Int
	FinalResult   *big.Int
	TxHash        string
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
	mathPromise := http.SendRequest(config, runtime, client, fetchMathResult, cre.ConsensusMedianAggregation[*big.Int]())
	offchainValue, err := mathPromise.Await()
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully fetched offchain value", "result", offchainValue)

	// Step 2: Read onchain data using the binding for the Storage contract
	evmClient := &evm.Client{
		ChainSelector: evmConfig.ChainSelector,
	}

	storageAddress := common.HexToAddress(evmConfig.StorageAddress)

	storageContract, err := storage.NewStorage(evmClient, storageAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract instance: %w", err)
	}
	onchainValue, err := storageContract.GetMultiplier(runtime, big.NewInt(-3)).Await()
	if err != nil {
		return nil, fmt.Errorf("failed to read onchain value: %w", err)
	}
	logger.Info("Successfully read onchain value", "result", onchainValue)

	// Step 3: Calculate the final result
	finalResultInt := new(big.Int).Add(onchainValue, offchainValue)

	logger.Info("Final calculated result", "result", finalResultInt)

	// Step 4: Write the result to the consumer contract
	txHash, err := updateCalculatorResult(config, runtime, evmConfig, offchainValue, onchainValue, finalResultInt)
	if err != nil {
		return nil, fmt.Errorf("failed to update calculator result: %w", err)
	}

	// Step 5: Log and return the final, consolidated result.
	finalWorkflowResult := &MyResult{
		OffchainValue: offchainValue,
		OnchainValue:  onchainValue,
		FinalResult:   finalResultInt,
		TxHash:        txHash,
	}

	logger.Info("Workflow finished successfully!", "result", finalWorkflowResult)

	return finalWorkflowResult, nil
}

func fetchMathResult(config *Config, logger *slog.Logger, sendRequester *http.SendRequester) (*big.Int, error) {
	req := &http.Request{Url: config.ApiUrl, Method: "GET"}
	resp, err := sendRequester.SendRequest(req).Await()
	if err != nil {
		return nil, fmt.Errorf("failed to get API response: %w", err)
	}

	// Parse the JSON response to extract ethereum.usd value
	var result struct {
		Ethereum struct {
			USD float64 `json:"usd"`
		} `json:"ethereum"`
	}

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Convert the float value to big.Int (multiply by 10^18 to preserve precision)
	// Adjust the multiplier based on your precision requirements
	val := new(big.Float).SetFloat64(result.Ethereum.USD)
	val.Mul(val, big.NewFloat(1e18))
	bigIntVal, _ := val.Int(nil)

	return bigIntVal, nil
}

// updateCalculatorResult handles the logic for writing data to the CalculatorConsumer contract.
func updateCalculatorResult(config *Config, runtime cre.Runtime, evmConfig EvmConfig, offchainValue *big.Int, onchainValue *big.Int, finalResult *big.Int) (string, error) {
	logger := runtime.Logger()
	logger.Info("Updating calculator result", "consumerAddress", evmConfig.CalculatorConsumerAddress)

	evmClient := &evm.Client{
		ChainSelector: evmConfig.ChainSelector,
	}

	// Create a contract binding instance pointed at the CalculatorConsumer address.
	consumerAddress := common.HexToAddress(evmConfig.CalculatorConsumerAddress)

	consumerContract, err := storage.NewStorage(evmClient, consumerAddress, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create consumer contract instance: %w", err)
	}

	gasConfig := &evm.GasConfig{
		GasLimit: evmConfig.GasLimit,
	}

	logger.Info("Writing to consumer contract", "finalResult", finalResult)

	// Encode the updateMultiplier call data
	updateInput := storage.UpdateMultiplierInput{NewETHPrice: finalResult}
	calldata, err := consumerContract.Codec.EncodeUpdateMultiplierMethodCall(updateInput)
	if err != nil {
		return "", fmt.Errorf("failed to encode updateMultiplier call: %w", err)
	}

	// Generate a signed CRE report containing the encoded data
	logger.Info("Generating CRE report")
	reportRequest := &cre.ReportRequest{
		EncodedPayload: calldata,
		EncoderName:    "evm",
		SigningAlgo:    "secp256k1",
		HashingAlgo:    "keccak256",
	}

	reportPromise := runtime.GenerateReport(reportRequest)
	report, err := reportPromise.Await()
	if err != nil {
		return "", fmt.Errorf("failed to generate report: %w", err)
	}

	// Use the WriteReport method to send the signed report to the consumer contract
	logger.Info("Sending report to consumer contract")
	writePromise := consumerContract.WriteReport(runtime, report, gasConfig)

	resp, err := writePromise.Await()
	if err != nil {
		logger.Error("WriteReport failed", "error", err)
		return "", fmt.Errorf("failed to write report: %w", err)
	}
	fmt.Println(resp, "<-----------------")
	txHash := fmt.Sprintf("0x%x", resp.TxHash)
	logger.Info("Report written successfully", "txHash", txHash)
	return txHash, nil
}

func main() {
	wasm.NewRunner(cre.ParseJSON[Config]).Run(InitWorkflow)
}
