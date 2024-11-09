// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ExecutionVerifierExecution is an auto generated low-level Go binding around an user-defined struct.
type ExecutionVerifierExecution struct {
	User       common.Address
	Nonce      *big.Int
	SellTokens []ExecutionVerifierTokenDetail
	BuyTokens  []ExecutionVerifierTokenDetail
	Routes     []ExecutionVerifierRoute
}

// ExecutionVerifierRoute is an auto generated low-level Go binding around an user-defined struct.
type ExecutionVerifierRoute struct {
	ProcessorIndex   *big.Int
	SellTokenDetails []ExecutionVerifierTokenDetail
	Payload          []byte
	Metadata         []byte
}

// ExecutionVerifierTokenDetail is an auto generated low-level Go binding around an user-defined struct.
type ExecutionVerifierTokenDetail struct {
	Token  common.Address
	Amount *big.Int
}

// AggregatorMetaData contains all meta data concerning the Aggregator contract.
var AggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_blacklister\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"BlacklistedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"presumedSigner\",\"type\":\"address\"}],\"name\":\"IncorrectSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InvalidEndState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"processorIndex\",\"type\":\"uint256\"}],\"name\":\"InvalidProcessor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenEntry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MathOverflowedMulDiv\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"MinBuyTokenAssertFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"UnauthorizedBlacklister\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"UnauthorizedPauser\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"sellTokens\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"buyTokens\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"processorIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"sellTokenDetails\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structExecutionVerifier.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"blacklister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"sellTokens\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"buyTokens\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"processorIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"sellTokenDetails\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExecutionVerifier.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"sellTokens\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"buyTokens\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"processorIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"sellTokenDetails\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExecutionVerifier.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"sellTokens\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"buyTokens\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"processorIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"sellTokenDetails\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExecutionVerifier.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"}],\"internalType\":\"structExecutionVerifier.Execution\",\"name\":\"execution\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"extract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeProportion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"sellTokens\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"buyTokens\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"processorIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structExecutionVerifier.TokenDetail[]\",\"name\":\"sellTokenDetails\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExecutionVerifier.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"}],\"internalType\":\"structExecutionVerifier.Execution\",\"name\":\"execution\",\"type\":\"tuple\"}],\"name\":\"hashExecution\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeProportion\",\"type\":\"uint256\"}],\"name\":\"setFeeProportion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIProcessor\",\"name\":\"processor\",\"type\":\"address\"}],\"name\":\"setProcessor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isBlacklisted\",\"type\":\"bool\"}],\"name\":\"setUserBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregatorMetaData.ABI instead.
var AggregatorABI = AggregatorMetaData.ABI

// Aggregator is an auto generated Go binding around an Ethereum contract.
type Aggregator struct {
	AggregatorCaller     // Read-only binding to the contract
	AggregatorTransactor // Write-only binding to the contract
	AggregatorFilterer   // Log filterer for contract events
}

// AggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregatorSession struct {
	Contract     *Aggregator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregatorCallerSession struct {
	Contract *AggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregatorTransactorSession struct {
	Contract     *AggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggregatorRaw struct {
	Contract *Aggregator // Generic contract binding to access the raw methods on
}

// AggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregatorCallerRaw struct {
	Contract *AggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// AggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregatorTransactorRaw struct {
	Contract *AggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregator creates a new instance of Aggregator, bound to a specific deployed contract.
func NewAggregator(address common.Address, backend bind.ContractBackend) (*Aggregator, error) {
	contract, err := bindAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aggregator{AggregatorCaller: AggregatorCaller{contract: contract}, AggregatorTransactor: AggregatorTransactor{contract: contract}, AggregatorFilterer: AggregatorFilterer{contract: contract}}, nil
}

// NewAggregatorCaller creates a new read-only instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorCaller(address common.Address, caller bind.ContractCaller) (*AggregatorCaller, error) {
	contract, err := bindAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorCaller{contract: contract}, nil
}

// NewAggregatorTransactor creates a new write-only instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregatorTransactor, error) {
	contract, err := bindAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorTransactor{contract: contract}, nil
}

// NewAggregatorFilterer creates a new log filterer instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregatorFilterer, error) {
	contract, err := bindAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatorFilterer{contract: contract}, nil
}

// bindAggregator binds a generic wrapper to an already deployed contract.
func bindAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggregator *AggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggregator.Contract.AggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggregator *AggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.Contract.AggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggregator *AggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggregator.Contract.AggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggregator *AggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggregator *AggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggregator *AggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggregator.Contract.contract.Transact(opts, method, params...)
}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Aggregator *AggregatorCaller) Blacklister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "blacklister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Aggregator *AggregatorSession) Blacklister() (common.Address, error) {
	return _Aggregator.Contract.Blacklister(&_Aggregator.CallOpts)
}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Aggregator *AggregatorCallerSession) Blacklister() (common.Address, error) {
	return _Aggregator.Contract.Blacklister(&_Aggregator.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Aggregator *AggregatorCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Aggregator *AggregatorSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Aggregator.Contract.Eip712Domain(&_Aggregator.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Aggregator *AggregatorCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Aggregator.Contract.Eip712Domain(&_Aggregator.CallOpts)
}

// Extract is a free data retrieval call binding the contract method 0x4a6ed8a2.
//
// Solidity: function extract((address,uint256,(address,uint256)[],(address,uint256)[],(uint256,(address,uint256)[],bytes,bytes)[]) execution, bytes signature) view returns(address)
func (_Aggregator *AggregatorCaller) Extract(opts *bind.CallOpts, execution ExecutionVerifierExecution, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "extract", execution, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Extract is a free data retrieval call binding the contract method 0x4a6ed8a2.
//
// Solidity: function extract((address,uint256,(address,uint256)[],(address,uint256)[],(uint256,(address,uint256)[],bytes,bytes)[]) execution, bytes signature) view returns(address)
func (_Aggregator *AggregatorSession) Extract(execution ExecutionVerifierExecution, signature []byte) (common.Address, error) {
	return _Aggregator.Contract.Extract(&_Aggregator.CallOpts, execution, signature)
}

// Extract is a free data retrieval call binding the contract method 0x4a6ed8a2.
//
// Solidity: function extract((address,uint256,(address,uint256)[],(address,uint256)[],(uint256,(address,uint256)[],bytes,bytes)[]) execution, bytes signature) view returns(address)
func (_Aggregator *AggregatorCallerSession) Extract(execution ExecutionVerifierExecution, signature []byte) (common.Address, error) {
	return _Aggregator.Contract.Extract(&_Aggregator.CallOpts, execution, signature)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Aggregator *AggregatorCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Aggregator *AggregatorSession) FeeCollector() (common.Address, error) {
	return _Aggregator.Contract.FeeCollector(&_Aggregator.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Aggregator *AggregatorCallerSession) FeeCollector() (common.Address, error) {
	return _Aggregator.Contract.FeeCollector(&_Aggregator.CallOpts)
}

// FeeProportion is a free data retrieval call binding the contract method 0x2a3ecc41.
//
// Solidity: function feeProportion() view returns(uint256)
func (_Aggregator *AggregatorCaller) FeeProportion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "feeProportion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeProportion is a free data retrieval call binding the contract method 0x2a3ecc41.
//
// Solidity: function feeProportion() view returns(uint256)
func (_Aggregator *AggregatorSession) FeeProportion() (*big.Int, error) {
	return _Aggregator.Contract.FeeProportion(&_Aggregator.CallOpts)
}

// FeeProportion is a free data retrieval call binding the contract method 0x2a3ecc41.
//
// Solidity: function feeProportion() view returns(uint256)
func (_Aggregator *AggregatorCallerSession) FeeProportion() (*big.Int, error) {
	return _Aggregator.Contract.FeeProportion(&_Aggregator.CallOpts)
}

// HashExecution is a free data retrieval call binding the contract method 0xcb69ec2f.
//
// Solidity: function hashExecution((address,uint256,(address,uint256)[],(address,uint256)[],(uint256,(address,uint256)[],bytes,bytes)[]) execution) view returns(bytes32)
func (_Aggregator *AggregatorCaller) HashExecution(opts *bind.CallOpts, execution ExecutionVerifierExecution) ([32]byte, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "hashExecution", execution)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashExecution is a free data retrieval call binding the contract method 0xcb69ec2f.
//
// Solidity: function hashExecution((address,uint256,(address,uint256)[],(address,uint256)[],(uint256,(address,uint256)[],bytes,bytes)[]) execution) view returns(bytes32)
func (_Aggregator *AggregatorSession) HashExecution(execution ExecutionVerifierExecution) ([32]byte, error) {
	return _Aggregator.Contract.HashExecution(&_Aggregator.CallOpts, execution)
}

// HashExecution is a free data retrieval call binding the contract method 0xcb69ec2f.
//
// Solidity: function hashExecution((address,uint256,(address,uint256)[],(address,uint256)[],(uint256,(address,uint256)[],bytes,bytes)[]) execution) view returns(bytes32)
func (_Aggregator *AggregatorCallerSession) HashExecution(execution ExecutionVerifierExecution) ([32]byte, error) {
	return _Aggregator.Contract.HashExecution(&_Aggregator.CallOpts, execution)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Aggregator *AggregatorCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Aggregator *AggregatorSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Aggregator.Contract.Nonces(&_Aggregator.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Aggregator *AggregatorCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Aggregator.Contract.Nonces(&_Aggregator.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggregator *AggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggregator *AggregatorSession) Owner() (common.Address, error) {
	return _Aggregator.Contract.Owner(&_Aggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggregator *AggregatorCallerSession) Owner() (common.Address, error) {
	return _Aggregator.Contract.Owner(&_Aggregator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Aggregator *AggregatorCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Aggregator *AggregatorSession) Paused() (bool, error) {
	return _Aggregator.Contract.Paused(&_Aggregator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Aggregator *AggregatorCallerSession) Paused() (bool, error) {
	return _Aggregator.Contract.Paused(&_Aggregator.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Aggregator *AggregatorCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Aggregator *AggregatorSession) Pauser() (common.Address, error) {
	return _Aggregator.Contract.Pauser(&_Aggregator.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Aggregator *AggregatorCallerSession) Pauser() (common.Address, error) {
	return _Aggregator.Contract.Pauser(&_Aggregator.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x8f73f125.
//
// Solidity: function execute((address,uint256)[] sellTokens, (address,uint256)[] buyTokens, (uint256,(address,uint256)[],bytes,bytes)[] routes) payable returns()
func (_Aggregator *AggregatorTransactor) Execute(opts *bind.TransactOpts, sellTokens []ExecutionVerifierTokenDetail, buyTokens []ExecutionVerifierTokenDetail, routes []ExecutionVerifierRoute) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "execute", sellTokens, buyTokens, routes)
}

// Execute is a paid mutator transaction binding the contract method 0x8f73f125.
//
// Solidity: function execute((address,uint256)[] sellTokens, (address,uint256)[] buyTokens, (uint256,(address,uint256)[],bytes,bytes)[] routes) payable returns()
func (_Aggregator *AggregatorSession) Execute(sellTokens []ExecutionVerifierTokenDetail, buyTokens []ExecutionVerifierTokenDetail, routes []ExecutionVerifierRoute) (*types.Transaction, error) {
	return _Aggregator.Contract.Execute(&_Aggregator.TransactOpts, sellTokens, buyTokens, routes)
}

// Execute is a paid mutator transaction binding the contract method 0x8f73f125.
//
// Solidity: function execute((address,uint256)[] sellTokens, (address,uint256)[] buyTokens, (uint256,(address,uint256)[],bytes,bytes)[] routes) payable returns()
func (_Aggregator *AggregatorTransactorSession) Execute(sellTokens []ExecutionVerifierTokenDetail, buyTokens []ExecutionVerifierTokenDetail, routes []ExecutionVerifierRoute) (*types.Transaction, error) {
	return _Aggregator.Contract.Execute(&_Aggregator.TransactOpts, sellTokens, buyTokens, routes)
}

// Execute0 is a paid mutator transaction binding the contract method 0xd6c04e5f.
//
// Solidity: function execute(address user, (address,uint256)[] sellTokens, (address,uint256)[] buyTokens, (uint256,(address,uint256)[],bytes,bytes)[] routes, bytes signature) payable returns()
func (_Aggregator *AggregatorTransactor) Execute0(opts *bind.TransactOpts, user common.Address, sellTokens []ExecutionVerifierTokenDetail, buyTokens []ExecutionVerifierTokenDetail, routes []ExecutionVerifierRoute, signature []byte) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "execute0", user, sellTokens, buyTokens, routes, signature)
}

// Execute0 is a paid mutator transaction binding the contract method 0xd6c04e5f.
//
// Solidity: function execute(address user, (address,uint256)[] sellTokens, (address,uint256)[] buyTokens, (uint256,(address,uint256)[],bytes,bytes)[] routes, bytes signature) payable returns()
func (_Aggregator *AggregatorSession) Execute0(user common.Address, sellTokens []ExecutionVerifierTokenDetail, buyTokens []ExecutionVerifierTokenDetail, routes []ExecutionVerifierRoute, signature []byte) (*types.Transaction, error) {
	return _Aggregator.Contract.Execute0(&_Aggregator.TransactOpts, user, sellTokens, buyTokens, routes, signature)
}

// Execute0 is a paid mutator transaction binding the contract method 0xd6c04e5f.
//
// Solidity: function execute(address user, (address,uint256)[] sellTokens, (address,uint256)[] buyTokens, (uint256,(address,uint256)[],bytes,bytes)[] routes, bytes signature) payable returns()
func (_Aggregator *AggregatorTransactorSession) Execute0(user common.Address, sellTokens []ExecutionVerifierTokenDetail, buyTokens []ExecutionVerifierTokenDetail, routes []ExecutionVerifierRoute, signature []byte) (*types.Transaction, error) {
	return _Aggregator.Contract.Execute0(&_Aggregator.TransactOpts, user, sellTokens, buyTokens, routes, signature)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Aggregator *AggregatorTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Aggregator *AggregatorSession) Pause() (*types.Transaction, error) {
	return _Aggregator.Contract.Pause(&_Aggregator.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Aggregator *AggregatorTransactorSession) Pause() (*types.Transaction, error) {
	return _Aggregator.Contract.Pause(&_Aggregator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggregator *AggregatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggregator *AggregatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aggregator.Contract.RenounceOwnership(&_Aggregator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggregator *AggregatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aggregator.Contract.RenounceOwnership(&_Aggregator.TransactOpts)
}

// SetFeeProportion is a paid mutator transaction binding the contract method 0x8e16ac45.
//
// Solidity: function setFeeProportion(uint256 _feeProportion) returns()
func (_Aggregator *AggregatorTransactor) SetFeeProportion(opts *bind.TransactOpts, _feeProportion *big.Int) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "setFeeProportion", _feeProportion)
}

// SetFeeProportion is a paid mutator transaction binding the contract method 0x8e16ac45.
//
// Solidity: function setFeeProportion(uint256 _feeProportion) returns()
func (_Aggregator *AggregatorSession) SetFeeProportion(_feeProportion *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.SetFeeProportion(&_Aggregator.TransactOpts, _feeProportion)
}

// SetFeeProportion is a paid mutator transaction binding the contract method 0x8e16ac45.
//
// Solidity: function setFeeProportion(uint256 _feeProportion) returns()
func (_Aggregator *AggregatorTransactorSession) SetFeeProportion(_feeProportion *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.SetFeeProportion(&_Aggregator.TransactOpts, _feeProportion)
}

// SetProcessor is a paid mutator transaction binding the contract method 0x15898b2e.
//
// Solidity: function setProcessor(uint256 index, address processor) returns()
func (_Aggregator *AggregatorTransactor) SetProcessor(opts *bind.TransactOpts, index *big.Int, processor common.Address) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "setProcessor", index, processor)
}

// SetProcessor is a paid mutator transaction binding the contract method 0x15898b2e.
//
// Solidity: function setProcessor(uint256 index, address processor) returns()
func (_Aggregator *AggregatorSession) SetProcessor(index *big.Int, processor common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.SetProcessor(&_Aggregator.TransactOpts, index, processor)
}

// SetProcessor is a paid mutator transaction binding the contract method 0x15898b2e.
//
// Solidity: function setProcessor(uint256 index, address processor) returns()
func (_Aggregator *AggregatorTransactorSession) SetProcessor(index *big.Int, processor common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.SetProcessor(&_Aggregator.TransactOpts, index, processor)
}

// SetUserBlacklist is a paid mutator transaction binding the contract method 0xccddf0b0.
//
// Solidity: function setUserBlacklist(address user, bool isBlacklisted) returns()
func (_Aggregator *AggregatorTransactor) SetUserBlacklist(opts *bind.TransactOpts, user common.Address, isBlacklisted bool) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "setUserBlacklist", user, isBlacklisted)
}

// SetUserBlacklist is a paid mutator transaction binding the contract method 0xccddf0b0.
//
// Solidity: function setUserBlacklist(address user, bool isBlacklisted) returns()
func (_Aggregator *AggregatorSession) SetUserBlacklist(user common.Address, isBlacklisted bool) (*types.Transaction, error) {
	return _Aggregator.Contract.SetUserBlacklist(&_Aggregator.TransactOpts, user, isBlacklisted)
}

// SetUserBlacklist is a paid mutator transaction binding the contract method 0xccddf0b0.
//
// Solidity: function setUserBlacklist(address user, bool isBlacklisted) returns()
func (_Aggregator *AggregatorTransactorSession) SetUserBlacklist(user common.Address, isBlacklisted bool) (*types.Transaction, error) {
	return _Aggregator.Contract.SetUserBlacklist(&_Aggregator.TransactOpts, user, isBlacklisted)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggregator *AggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggregator *AggregatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.TransferOwnership(&_Aggregator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggregator *AggregatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.TransferOwnership(&_Aggregator.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Aggregator *AggregatorTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Aggregator *AggregatorSession) Unpause() (*types.Transaction, error) {
	return _Aggregator.Contract.Unpause(&_Aggregator.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Aggregator *AggregatorTransactorSession) Unpause() (*types.Transaction, error) {
	return _Aggregator.Contract.Unpause(&_Aggregator.TransactOpts)
}

// AggregatorEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Aggregator contract.
type AggregatorEIP712DomainChangedIterator struct {
	Event *AggregatorEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AggregatorEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AggregatorEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AggregatorEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorEIP712DomainChanged represents a EIP712DomainChanged event raised by the Aggregator contract.
type AggregatorEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Aggregator *AggregatorFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*AggregatorEIP712DomainChangedIterator, error) {

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &AggregatorEIP712DomainChangedIterator{contract: _Aggregator.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Aggregator *AggregatorFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *AggregatorEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorEIP712DomainChanged)
				if err := _Aggregator.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Aggregator *AggregatorFilterer) ParseEIP712DomainChanged(log types.Log) (*AggregatorEIP712DomainChanged, error) {
	event := new(AggregatorEIP712DomainChanged)
	if err := _Aggregator.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the Aggregator contract.
type AggregatorExecutedIterator struct {
	Event *AggregatorExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AggregatorExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AggregatorExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AggregatorExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorExecuted represents a Executed event raised by the Aggregator contract.
type AggregatorExecuted struct {
	User       common.Address
	SellTokens []ExecutionVerifierTokenDetail
	BuyTokens  []ExecutionVerifierTokenDetail
	Routes     []ExecutionVerifierRoute
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x86b324412f1ae3de8b6f67889e726a5e6de104beb3ddc456dbff9f9dff8964bb.
//
// Solidity: event Executed(address indexed user, (address,uint256)[] sellTokens, (address,uint256)[] buyTokens, (uint256,(address,uint256)[],bytes,bytes)[] routes)
func (_Aggregator *AggregatorFilterer) FilterExecuted(opts *bind.FilterOpts, user []common.Address) (*AggregatorExecutedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "Executed", userRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorExecutedIterator{contract: _Aggregator.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x86b324412f1ae3de8b6f67889e726a5e6de104beb3ddc456dbff9f9dff8964bb.
//
// Solidity: event Executed(address indexed user, (address,uint256)[] sellTokens, (address,uint256)[] buyTokens, (uint256,(address,uint256)[],bytes,bytes)[] routes)
func (_Aggregator *AggregatorFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *AggregatorExecuted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "Executed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorExecuted)
				if err := _Aggregator.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecuted is a log parse operation binding the contract event 0x86b324412f1ae3de8b6f67889e726a5e6de104beb3ddc456dbff9f9dff8964bb.
//
// Solidity: event Executed(address indexed user, (address,uint256)[] sellTokens, (address,uint256)[] buyTokens, (uint256,(address,uint256)[],bytes,bytes)[] routes)
func (_Aggregator *AggregatorFilterer) ParseExecuted(log types.Log) (*AggregatorExecuted, error) {
	event := new(AggregatorExecuted)
	if err := _Aggregator.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Aggregator contract.
type AggregatorOwnershipTransferredIterator struct {
	Event *AggregatorOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AggregatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AggregatorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorOwnershipTransferred represents a OwnershipTransferred event raised by the Aggregator contract.
type AggregatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggregator *AggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AggregatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorOwnershipTransferredIterator{contract: _Aggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggregator *AggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AggregatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorOwnershipTransferred)
				if err := _Aggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggregator *AggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*AggregatorOwnershipTransferred, error) {
	event := new(AggregatorOwnershipTransferred)
	if err := _Aggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Aggregator contract.
type AggregatorPausedIterator struct {
	Event *AggregatorPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AggregatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AggregatorPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AggregatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorPaused represents a Paused event raised by the Aggregator contract.
type AggregatorPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Aggregator *AggregatorFilterer) FilterPaused(opts *bind.FilterOpts) (*AggregatorPausedIterator, error) {

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AggregatorPausedIterator{contract: _Aggregator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Aggregator *AggregatorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AggregatorPaused) (event.Subscription, error) {

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorPaused)
				if err := _Aggregator.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Aggregator *AggregatorFilterer) ParsePaused(log types.Log) (*AggregatorPaused, error) {
	event := new(AggregatorPaused)
	if err := _Aggregator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Aggregator contract.
type AggregatorUnpausedIterator struct {
	Event *AggregatorUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AggregatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AggregatorUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AggregatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorUnpaused represents a Unpaused event raised by the Aggregator contract.
type AggregatorUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Aggregator *AggregatorFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AggregatorUnpausedIterator, error) {

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AggregatorUnpausedIterator{contract: _Aggregator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Aggregator *AggregatorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AggregatorUnpaused) (event.Subscription, error) {

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorUnpaused)
				if err := _Aggregator.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Aggregator *AggregatorFilterer) ParseUnpaused(log types.Log) (*AggregatorUnpaused, error) {
	event := new(AggregatorUnpaused)
	if err := _Aggregator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
