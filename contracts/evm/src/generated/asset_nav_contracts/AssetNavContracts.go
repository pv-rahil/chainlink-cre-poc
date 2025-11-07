// Code generated — DO NOT EDIT.

package asset_nav_contracts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb2 "github.com/smartcontractkit/chainlink-protos/cre/go/sdk"
	"github.com/smartcontractkit/chainlink-protos/cre/go/values/pb"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Sprintf
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = emptypb.Empty{}
	_ = pb.NewBigIntFromInt
	_ = pb2.AggregationType_AGGREGATION_TYPE_COMMON_PREFIX
	_ = bindings.FilterOptions{}
	_ = evm.FilterLogTriggerRequest{}
	_ = cre.ResponseBufferTooSmall
	_ = rpc.API{}
	_ = json.Unmarshal
)

var AssetNavContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expectedAuthor\",\"type\":\"address\"},{\"internalType\":\"bytes10\",\"name\":\"expectedWorkflowName\",\"type\":\"bytes10\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"received\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"InvalidAuthor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes10\",\"name\":\"received\",\"type\":\"bytes10\"},{\"internalType\":\"bytes10\",\"name\":\"expected\",\"type\":\"bytes10\"}],\"name\":\"InvalidWorkflowName\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"newAssetNav\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AssetNavUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EXPECTED_AUTHOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXPECTED_WORKFLOW_NAME\",\"outputs\":[{\"internalType\":\"bytes10\",\"name\":\"\",\"type\":\"bytes10\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetNAV\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNAVAtIndex\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"nav\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNAVHistoryLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"navHistory\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"nav\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"onReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// Structs

// Contract Method Inputs
type GetNAVAtIndexInput struct {
	Index *big.Int
}

type NavHistoryInput struct {
	Arg0 *big.Int
}

type OnReportInput struct {
	Arg0   []byte
	Report []byte
}

type SupportsInterfaceInput struct {
	InterfaceId [4]byte
}

// Contract Method Outputs
type GetNAVAtIndexOutput struct {
	Nav       *big.Int
	Timestamp *big.Int
}

type NavHistoryOutput struct {
	Nav       *big.Int
	Timestamp *big.Int
}

// Errors
type InvalidAuthor struct {
	Received common.Address
	Expected common.Address
}

type InvalidWorkflowName struct {
	Received [10]byte
	Expected [10]byte
}

// Events
type AssetNavUpdated struct {
	NewAssetNav *big.Int
	Timestamp   *big.Int
}

// Main Binding Type for AssetNavContracts
type AssetNavContracts struct {
	Address common.Address
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   AssetNavContractsCodec
}

type AssetNavContractsCodec interface {
	EncodeEXPECTEDAUTHORMethodCall() ([]byte, error)
	DecodeEXPECTEDAUTHORMethodOutput(data []byte) (common.Address, error)
	EncodeEXPECTEDWORKFLOWNAMEMethodCall() ([]byte, error)
	DecodeEXPECTEDWORKFLOWNAMEMethodOutput(data []byte) ([10]byte, error)
	EncodeAssetNAVMethodCall() ([]byte, error)
	DecodeAssetNAVMethodOutput(data []byte) (*big.Int, error)
	EncodeGetNAVAtIndexMethodCall(in GetNAVAtIndexInput) ([]byte, error)
	DecodeGetNAVAtIndexMethodOutput(data []byte) (GetNAVAtIndexOutput, error)
	EncodeGetNAVHistoryLengthMethodCall() ([]byte, error)
	DecodeGetNAVHistoryLengthMethodOutput(data []byte) (*big.Int, error)
	EncodeNavHistoryMethodCall(in NavHistoryInput) ([]byte, error)
	DecodeNavHistoryMethodOutput(data []byte) (NavHistoryOutput, error)
	EncodeOnReportMethodCall(in OnReportInput) ([]byte, error)
	EncodeSupportsInterfaceMethodCall(in SupportsInterfaceInput) ([]byte, error)
	DecodeSupportsInterfaceMethodOutput(data []byte) (bool, error)
	AssetNavUpdatedLogHash() []byte
	EncodeAssetNavUpdatedTopics(evt abi.Event, values []AssetNavUpdated) ([]*evm.TopicValues, error)
	DecodeAssetNavUpdated(log *evm.Log) (*AssetNavUpdated, error)
}

func NewAssetNavContracts(
	client *evm.Client,
	address common.Address,
	options *bindings.ContractInitOptions,
) (*AssetNavContracts, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetNavContractsMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewCodec()
	if err != nil {
		return nil, err
	}
	return &AssetNavContracts{
		Address: address,
		Options: options,
		ABI:     &parsed,
		client:  client,
		Codec:   codec,
	}, nil
}

type Codec struct {
	abi *abi.ABI
}

func NewCodec() (AssetNavContractsCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetNavContractsMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &Codec{abi: &parsed}, nil
}

func (c *Codec) EncodeEXPECTEDAUTHORMethodCall() ([]byte, error) {
	return c.abi.Pack("EXPECTED_AUTHOR")
}

func (c *Codec) DecodeEXPECTEDAUTHORMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["EXPECTED_AUTHOR"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeEXPECTEDWORKFLOWNAMEMethodCall() ([]byte, error) {
	return c.abi.Pack("EXPECTED_WORKFLOW_NAME")
}

func (c *Codec) DecodeEXPECTEDWORKFLOWNAMEMethodOutput(data []byte) ([10]byte, error) {
	vals, err := c.abi.Methods["EXPECTED_WORKFLOW_NAME"].Outputs.Unpack(data)
	if err != nil {
		return *new([10]byte), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new([10]byte), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result [10]byte
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new([10]byte), fmt.Errorf("failed to unmarshal to [10]byte: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeAssetNAVMethodCall() ([]byte, error) {
	return c.abi.Pack("assetNAV")
}

func (c *Codec) DecodeAssetNAVMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["assetNAV"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetNAVAtIndexMethodCall(in GetNAVAtIndexInput) ([]byte, error) {
	return c.abi.Pack("getNAVAtIndex", in.Index)
}

func (c *Codec) DecodeGetNAVAtIndexMethodOutput(data []byte) (GetNAVAtIndexOutput, error) {
	vals, err := c.abi.Methods["getNAVAtIndex"].Outputs.Unpack(data)
	if err != nil {
		return GetNAVAtIndexOutput{}, err
	}
	if len(vals) != 2 {
		return GetNAVAtIndexOutput{}, fmt.Errorf("expected 2 values, got %d", len(vals))
	}
	jsonData0, err := json.Marshal(vals[0])
	if err != nil {
		return GetNAVAtIndexOutput{}, fmt.Errorf("failed to marshal ABI result 0: %w", err)
	}

	var result0 *big.Int
	if err := json.Unmarshal(jsonData0, &result0); err != nil {
		return GetNAVAtIndexOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData1, err := json.Marshal(vals[1])
	if err != nil {
		return GetNAVAtIndexOutput{}, fmt.Errorf("failed to marshal ABI result 1: %w", err)
	}

	var result1 *big.Int
	if err := json.Unmarshal(jsonData1, &result1); err != nil {
		return GetNAVAtIndexOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return GetNAVAtIndexOutput{
		Nav:       result0,
		Timestamp: result1,
	}, nil
}

func (c *Codec) EncodeGetNAVHistoryLengthMethodCall() ([]byte, error) {
	return c.abi.Pack("getNAVHistoryLength")
}

func (c *Codec) DecodeGetNAVHistoryLengthMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["getNAVHistoryLength"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeNavHistoryMethodCall(in NavHistoryInput) ([]byte, error) {
	return c.abi.Pack("navHistory", in.Arg0)
}

func (c *Codec) DecodeNavHistoryMethodOutput(data []byte) (NavHistoryOutput, error) {
	vals, err := c.abi.Methods["navHistory"].Outputs.Unpack(data)
	if err != nil {
		return NavHistoryOutput{}, err
	}
	if len(vals) != 2 {
		return NavHistoryOutput{}, fmt.Errorf("expected 2 values, got %d", len(vals))
	}
	jsonData0, err := json.Marshal(vals[0])
	if err != nil {
		return NavHistoryOutput{}, fmt.Errorf("failed to marshal ABI result 0: %w", err)
	}

	var result0 *big.Int
	if err := json.Unmarshal(jsonData0, &result0); err != nil {
		return NavHistoryOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData1, err := json.Marshal(vals[1])
	if err != nil {
		return NavHistoryOutput{}, fmt.Errorf("failed to marshal ABI result 1: %w", err)
	}

	var result1 *big.Int
	if err := json.Unmarshal(jsonData1, &result1); err != nil {
		return NavHistoryOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return NavHistoryOutput{
		Nav:       result0,
		Timestamp: result1,
	}, nil
}

func (c *Codec) EncodeOnReportMethodCall(in OnReportInput) ([]byte, error) {
	return c.abi.Pack("onReport", in.Arg0, in.Report)
}

func (c *Codec) EncodeSupportsInterfaceMethodCall(in SupportsInterfaceInput) ([]byte, error) {
	return c.abi.Pack("supportsInterface", in.InterfaceId)
}

func (c *Codec) DecodeSupportsInterfaceMethodOutput(data []byte) (bool, error) {
	vals, err := c.abi.Methods["supportsInterface"].Outputs.Unpack(data)
	if err != nil {
		return *new(bool), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(bool), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result bool
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(bool), fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return result, nil
}

func (c *Codec) AssetNavUpdatedLogHash() []byte {
	return c.abi.Events["AssetNavUpdated"].ID.Bytes()
}

func (c *Codec) EncodeAssetNavUpdatedTopics(
	evt abi.Event,
	values []AssetNavUpdated,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	topics := make([]*evm.TopicValues, len(rawTopics)+1)
	topics[0] = &evm.TopicValues{
		Values: [][]byte{evt.ID.Bytes()},
	}
	for i, hashList := range rawTopics {
		bs := make([][]byte, len(hashList))
		for j, h := range hashList {
			bs[j] = h.Bytes()
		}
		topics[i+1] = &evm.TopicValues{Values: bs}
	}
	return topics, nil
}

// DecodeAssetNavUpdated decodes a log into a AssetNavUpdated struct.
func (c *Codec) DecodeAssetNavUpdated(log *evm.Log) (*AssetNavUpdated, error) {
	event := new(AssetNavUpdated)
	if err := c.abi.UnpackIntoInterface(event, "AssetNavUpdated", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["AssetNavUpdated"].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c AssetNavContracts) EXPECTEDAUTHOR(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeEXPECTEDAUTHORMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeEXPECTEDAUTHORMethodOutput(response.Data)
	})

}

func (c AssetNavContracts) EXPECTEDWORKFLOWNAME(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[[10]byte] {
	calldata, err := c.Codec.EncodeEXPECTEDWORKFLOWNAMEMethodCall()
	if err != nil {
		return cre.PromiseFromResult[[10]byte](*new([10]byte), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) ([10]byte, error) {
		return c.Codec.DecodeEXPECTEDWORKFLOWNAMEMethodOutput(response.Data)
	})

}

func (c AssetNavContracts) AssetNAV(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeAssetNAVMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*big.Int](*new(*big.Int), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (*big.Int, error) {
		return c.Codec.DecodeAssetNAVMethodOutput(response.Data)
	})

}

func (c AssetNavContracts) GetNAVAtIndex(
	runtime cre.Runtime,
	args GetNAVAtIndexInput,
	blockNumber *big.Int,
) cre.Promise[GetNAVAtIndexOutput] {
	calldata, err := c.Codec.EncodeGetNAVAtIndexMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[GetNAVAtIndexOutput](GetNAVAtIndexOutput{}, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (GetNAVAtIndexOutput, error) {
		return c.Codec.DecodeGetNAVAtIndexMethodOutput(response.Data)
	})

}

func (c AssetNavContracts) GetNAVHistoryLength(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeGetNAVHistoryLengthMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*big.Int](*new(*big.Int), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (*big.Int, error) {
		return c.Codec.DecodeGetNAVHistoryLengthMethodOutput(response.Data)
	})

}

func (c AssetNavContracts) NavHistory(
	runtime cre.Runtime,
	args NavHistoryInput,
	blockNumber *big.Int,
) cre.Promise[NavHistoryOutput] {
	calldata, err := c.Codec.EncodeNavHistoryMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[NavHistoryOutput](NavHistoryOutput{}, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (NavHistoryOutput, error) {
		return c.Codec.DecodeNavHistoryMethodOutput(response.Data)
	})

}

func (c AssetNavContracts) WriteReport(
	runtime cre.Runtime,
	report *cre.Report,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver:  c.Address.Bytes(),
		Report:    report,
		GasConfig: gasConfig,
	})
}

// DecodeInvalidAuthorError decodes a InvalidAuthor error from revert data.
func (c *AssetNavContracts) DecodeInvalidAuthorError(data []byte) (*InvalidAuthor, error) {
	args := c.ABI.Errors["InvalidAuthor"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 2 {
		return nil, fmt.Errorf("expected 2 values, got %d", len(values))
	}

	received, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for received in InvalidAuthor error")
	}

	expected, ok1 := values[1].(common.Address)
	if !ok1 {
		return nil, fmt.Errorf("unexpected type for expected in InvalidAuthor error")
	}

	return &InvalidAuthor{
		Received: received,
		Expected: expected,
	}, nil
}

// Error implements the error interface for InvalidAuthor.
func (e *InvalidAuthor) Error() string {
	return fmt.Sprintf("InvalidAuthor error: received=%v; expected=%v;", e.Received, e.Expected)
}

// DecodeInvalidWorkflowNameError decodes a InvalidWorkflowName error from revert data.
func (c *AssetNavContracts) DecodeInvalidWorkflowNameError(data []byte) (*InvalidWorkflowName, error) {
	args := c.ABI.Errors["InvalidWorkflowName"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 2 {
		return nil, fmt.Errorf("expected 2 values, got %d", len(values))
	}

	received, ok0 := values[0].([10]byte)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for received in InvalidWorkflowName error")
	}

	expected, ok1 := values[1].([10]byte)
	if !ok1 {
		return nil, fmt.Errorf("unexpected type for expected in InvalidWorkflowName error")
	}

	return &InvalidWorkflowName{
		Received: received,
		Expected: expected,
	}, nil
}

// Error implements the error interface for InvalidWorkflowName.
func (e *InvalidWorkflowName) Error() string {
	return fmt.Sprintf("InvalidWorkflowName error: received=%v; expected=%v;", e.Received, e.Expected)
}

func (c *AssetNavContracts) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	case common.Bytes2Hex(c.ABI.Errors["InvalidAuthor"].ID.Bytes()[:4]):
		return c.DecodeInvalidAuthorError(data)
	case common.Bytes2Hex(c.ABI.Errors["InvalidWorkflowName"].ID.Bytes()[:4]):
		return c.DecodeInvalidWorkflowNameError(data)
	default:
		return nil, errors.New("unknown error selector")
	}
}

func (c *AssetNavContracts) LogTriggerAssetNavUpdatedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []AssetNavUpdated) (cre.Trigger[*evm.Log, *evm.Log], error) {
	event := c.ABI.Events["AssetNavUpdated"]
	topics, err := c.Codec.EncodeAssetNavUpdatedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for AssetNavUpdated: %w", err)
	}

	return evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	}), nil
}

func (c *AssetNavContracts) FilterLogsAssetNavUpdated(runtime cre.Runtime, options *bindings.FilterOptions) cre.Promise[*evm.FilterLogsReply] {
	if options == nil {
		options = &bindings.FilterOptions{
			ToBlock: options.ToBlock,
		}
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.AssetNavUpdatedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	})
}
