// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package putonchain

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

// PutonchainMetaData contains all meta data concerning the Putonchain contract.
var PutonchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_price\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_label\",\"type\":\"string\"}],\"name\":\"addProject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getProject\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PutonchainABI is the input ABI used to generate the binding from.
// Deprecated: Use PutonchainMetaData.ABI instead.
var PutonchainABI = PutonchainMetaData.ABI

// Putonchain is an auto generated Go binding around an Ethereum contract.
type Putonchain struct {
	PutonchainCaller     // Read-only binding to the contract
	PutonchainTransactor // Write-only binding to the contract
	PutonchainFilterer   // Log filterer for contract events
}

// PutonchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type PutonchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PutonchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PutonchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PutonchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PutonchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PutonchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PutonchainSession struct {
	Contract     *Putonchain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PutonchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PutonchainCallerSession struct {
	Contract *PutonchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PutonchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PutonchainTransactorSession struct {
	Contract     *PutonchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PutonchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type PutonchainRaw struct {
	Contract *Putonchain // Generic contract binding to access the raw methods on
}

// PutonchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PutonchainCallerRaw struct {
	Contract *PutonchainCaller // Generic read-only contract binding to access the raw methods on
}

// PutonchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PutonchainTransactorRaw struct {
	Contract *PutonchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPutonchain creates a new instance of Putonchain, bound to a specific deployed contract.
func NewPutonchain(address common.Address, backend bind.ContractBackend) (*Putonchain, error) {
	contract, err := bindPutonchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Putonchain{PutonchainCaller: PutonchainCaller{contract: contract}, PutonchainTransactor: PutonchainTransactor{contract: contract}, PutonchainFilterer: PutonchainFilterer{contract: contract}}, nil
}

// NewPutonchainCaller creates a new read-only instance of Putonchain, bound to a specific deployed contract.
func NewPutonchainCaller(address common.Address, caller bind.ContractCaller) (*PutonchainCaller, error) {
	contract, err := bindPutonchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PutonchainCaller{contract: contract}, nil
}

// NewPutonchainTransactor creates a new write-only instance of Putonchain, bound to a specific deployed contract.
func NewPutonchainTransactor(address common.Address, transactor bind.ContractTransactor) (*PutonchainTransactor, error) {
	contract, err := bindPutonchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PutonchainTransactor{contract: contract}, nil
}

// NewPutonchainFilterer creates a new log filterer instance of Putonchain, bound to a specific deployed contract.
func NewPutonchainFilterer(address common.Address, filterer bind.ContractFilterer) (*PutonchainFilterer, error) {
	contract, err := bindPutonchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PutonchainFilterer{contract: contract}, nil
}

// bindPutonchain binds a generic wrapper to an already deployed contract.
func bindPutonchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PutonchainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Putonchain *PutonchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Putonchain.Contract.PutonchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Putonchain *PutonchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Putonchain.Contract.PutonchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Putonchain *PutonchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Putonchain.Contract.PutonchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Putonchain *PutonchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Putonchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Putonchain *PutonchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Putonchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Putonchain *PutonchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Putonchain.Contract.contract.Transact(opts, method, params...)
}

// GetProject is a free data retrieval call binding the contract method 0xf0f3f2c8.
//
// Solidity: function getProject(uint256 _id) view returns(uint256, string, string)
func (_Putonchain *PutonchainCaller) GetProject(opts *bind.CallOpts, _id *big.Int) (*big.Int, string, string, error) {
	var out []interface{}
	err := _Putonchain.contract.Call(opts, &out, "getProject", _id)

	if err != nil {
		return *new(*big.Int), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// GetProject is a free data retrieval call binding the contract method 0xf0f3f2c8.
//
// Solidity: function getProject(uint256 _id) view returns(uint256, string, string)
func (_Putonchain *PutonchainSession) GetProject(_id *big.Int) (*big.Int, string, string, error) {
	return _Putonchain.Contract.GetProject(&_Putonchain.CallOpts, _id)
}

// GetProject is a free data retrieval call binding the contract method 0xf0f3f2c8.
//
// Solidity: function getProject(uint256 _id) view returns(uint256, string, string)
func (_Putonchain *PutonchainCallerSession) GetProject(_id *big.Int) (*big.Int, string, string, error) {
	return _Putonchain.Contract.GetProject(&_Putonchain.CallOpts, _id)
}

// AddProject is a paid mutator transaction binding the contract method 0x849d0faa.
//
// Solidity: function addProject(uint256 _id, string _price, string _label) returns()
func (_Putonchain *PutonchainTransactor) AddProject(opts *bind.TransactOpts, _id *big.Int, _price string, _label string) (*types.Transaction, error) {
	return _Putonchain.contract.Transact(opts, "addProject", _id, _price, _label)
}

// AddProject is a paid mutator transaction binding the contract method 0x849d0faa.
//
// Solidity: function addProject(uint256 _id, string _price, string _label) returns()
func (_Putonchain *PutonchainSession) AddProject(_id *big.Int, _price string, _label string) (*types.Transaction, error) {
	return _Putonchain.Contract.AddProject(&_Putonchain.TransactOpts, _id, _price, _label)
}

// AddProject is a paid mutator transaction binding the contract method 0x849d0faa.
//
// Solidity: function addProject(uint256 _id, string _price, string _label) returns()
func (_Putonchain *PutonchainTransactorSession) AddProject(_id *big.Int, _price string, _label string) (*types.Transaction, error) {
	return _Putonchain.Contract.AddProject(&_Putonchain.TransactOpts, _id, _price, _label)
}
