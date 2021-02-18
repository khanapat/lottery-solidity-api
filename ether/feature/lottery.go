// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feature

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FeatureABI is the input ABI used to generate the binding from.
const FeatureABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Address\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkVoteByAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"countVoteNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"setJackpot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"voteTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FeatureBin is the compiled bytecode used for deploying new contracts.
var FeatureBin = "0x608060405234801561001057600080fd5b50604051610b2f380380610b2f8339818101604052602081101561003357600080fd5b810190808051906020019092919050505080600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600660036101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600660006101000a81548160ff0219169083151502179055506001600660026101000a81548160ff02191690831515021790555050610a238061010c6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063741f76cd1161005b578063741f76cd146100fc578063a22b425614610106578063be8b58c014610182578063efd0ad14146101c45761007d565b80633d7b995c1461008257806359992cc8146100b0578063722713f7146100de575b600080fd5b6100ae6004803603602081101561009857600080fd5b81019080803590602001909291905050506101e2565b005b6100dc600480360360208110156100c657600080fd5b810190808035906020019092919050505061044e565b005b6100e66105ee565b6040518082815260200191505060405180910390f35b6101046106b9565b005b61010e610929565b604051808373ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019060200280838360005b8381101561016d578082015181840152602081019050610152565b50505050905001935050505060405180910390f35b6101ae6004803603602081101561019857600080fd5b81019080803590602001909291905050506109c6565b6040518082815260200191505060405180910390f35b6101cc6109e3565b6040518082815260200191505060405180910390f35b600660009054906101000a900460ff16610264576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f566f746520697320636c6f73656400000000000000000000000000000000000081525060200191505060405180910390fd5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190806001815401808255809150506001900390600052602060002001600090919091909150556003600082815260200190815260200160002060008154809291906001019190505550600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082815260200190815260200160002060008154809291906001019190505550600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd333060646040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b1580156103fe57600080fd5b505af1158015610412573d6000803e3d6000fd5b505050506040513d602081101561042857600080fd5b810190808051906020019092919050505050606460046000828254019250508190555050565b600660039054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610511576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f596f75277265206e6f7420617574686f72697a6564000000000000000000000081525060200191505060405180910390fd5b600660029054906101000a900460ff16610593576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f4a61636b706f742068617320616c72656164792073657400000000000000000081525060200191505060405180910390fd5b806005819055506001600660016101000a81548160ff0219169083151502179055506000600660006101000a81548160ff0219169083151502179055506000600660026101000a81548160ff02191690831515021790555050565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561067957600080fd5b505afa15801561068d573d6000803e3d6000fd5b505050506040513d60208110156106a357600080fd5b8101908080519060200190929190505050905090565b600660019054906101000a900460ff1661073b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f4a61636b706f742773206e6f742073657400000000000000000000000000000081525060200191505060405180910390fd5b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006005548152602001908152602001600020541415610804576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f596f75277265206c6f736572000000000000000000000000000000000000000081525060200191505060405180910390fd5b6000600360006005548152602001908152602001600020546004548161082657fe5b049050600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3033846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b1580156108da57600080fd5b505af11580156108ee573d6000803e3d6000fd5b505050506040513d602081101561090457600080fd5b8101908080519060200190929190505050508060046000828254039250508190555050565b60006060336000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020808054806020026020016040519081016040528092919081815260200182805480156109b757602002820191906000526020600020905b8154815260200190600101908083116109a3575b50505050509050915091509091565b600060036000838152602001908152602001600020549050919050565b600060045490509056fea2646970667358221220952d9ecf6d3238d9f3c5b88a2768ac4562b9aad253bd4d6a70d57a76816082e164736f6c63430007030033"

// DeployFeature deploys a new Ethereum contract, binding an instance of Feature to it.
func DeployFeature(auth *bind.TransactOpts, backend bind.ContractBackend, erc20Address common.Address) (common.Address, *types.Transaction, *Feature, error) {
	parsed, err := abi.JSON(strings.NewReader(FeatureABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FeatureBin), backend, erc20Address)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Feature{FeatureCaller: FeatureCaller{contract: contract}, FeatureTransactor: FeatureTransactor{contract: contract}, FeatureFilterer: FeatureFilterer{contract: contract}}, nil
}

// Feature is an auto generated Go binding around an Ethereum contract.
type Feature struct {
	FeatureCaller     // Read-only binding to the contract
	FeatureTransactor // Write-only binding to the contract
	FeatureFilterer   // Log filterer for contract events
}

// FeatureCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeatureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeatureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeatureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeatureFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeatureFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeatureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeatureSession struct {
	Contract     *Feature          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeatureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeatureCallerSession struct {
	Contract *FeatureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FeatureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeatureTransactorSession struct {
	Contract     *FeatureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FeatureRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeatureRaw struct {
	Contract *Feature // Generic contract binding to access the raw methods on
}

// FeatureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeatureCallerRaw struct {
	Contract *FeatureCaller // Generic read-only contract binding to access the raw methods on
}

// FeatureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeatureTransactorRaw struct {
	Contract *FeatureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeature creates a new instance of Feature, bound to a specific deployed contract.
func NewFeature(address common.Address, backend bind.ContractBackend) (*Feature, error) {
	contract, err := bindFeature(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Feature{FeatureCaller: FeatureCaller{contract: contract}, FeatureTransactor: FeatureTransactor{contract: contract}, FeatureFilterer: FeatureFilterer{contract: contract}}, nil
}

// NewFeatureCaller creates a new read-only instance of Feature, bound to a specific deployed contract.
func NewFeatureCaller(address common.Address, caller bind.ContractCaller) (*FeatureCaller, error) {
	contract, err := bindFeature(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeatureCaller{contract: contract}, nil
}

// NewFeatureTransactor creates a new write-only instance of Feature, bound to a specific deployed contract.
func NewFeatureTransactor(address common.Address, transactor bind.ContractTransactor) (*FeatureTransactor, error) {
	contract, err := bindFeature(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeatureTransactor{contract: contract}, nil
}

// NewFeatureFilterer creates a new log filterer instance of Feature, bound to a specific deployed contract.
func NewFeatureFilterer(address common.Address, filterer bind.ContractFilterer) (*FeatureFilterer, error) {
	contract, err := bindFeature(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeatureFilterer{contract: contract}, nil
}

// bindFeature binds a generic wrapper to an already deployed contract.
func bindFeature(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeatureABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Feature *FeatureRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Feature.Contract.FeatureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Feature *FeatureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feature.Contract.FeatureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Feature *FeatureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Feature.Contract.FeatureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Feature *FeatureCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Feature.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Feature *FeatureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feature.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Feature *FeatureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Feature.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_Feature *FeatureCaller) BalanceOf(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Feature.contract.Call(opts, &out, "balanceOf")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_Feature *FeatureSession) BalanceOf() (*big.Int, error) {
	return _Feature.Contract.BalanceOf(&_Feature.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_Feature *FeatureCallerSession) BalanceOf() (*big.Int, error) {
	return _Feature.Contract.BalanceOf(&_Feature.CallOpts)
}

// CheckTotal is a free data retrieval call binding the contract method 0xefd0ad14.
//
// Solidity: function checkTotal() view returns(uint256)
func (_Feature *FeatureCaller) CheckTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Feature.contract.Call(opts, &out, "checkTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckTotal is a free data retrieval call binding the contract method 0xefd0ad14.
//
// Solidity: function checkTotal() view returns(uint256)
func (_Feature *FeatureSession) CheckTotal() (*big.Int, error) {
	return _Feature.Contract.CheckTotal(&_Feature.CallOpts)
}

// CheckTotal is a free data retrieval call binding the contract method 0xefd0ad14.
//
// Solidity: function checkTotal() view returns(uint256)
func (_Feature *FeatureCallerSession) CheckTotal() (*big.Int, error) {
	return _Feature.Contract.CheckTotal(&_Feature.CallOpts)
}

// CheckVoteByAddress is a free data retrieval call binding the contract method 0xa22b4256.
//
// Solidity: function checkVoteByAddress() view returns(address, uint256[])
func (_Feature *FeatureCaller) CheckVoteByAddress(opts *bind.CallOpts) (common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Feature.contract.Call(opts, &out, "checkVoteByAddress")

	if err != nil {
		return *new(common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// CheckVoteByAddress is a free data retrieval call binding the contract method 0xa22b4256.
//
// Solidity: function checkVoteByAddress() view returns(address, uint256[])
func (_Feature *FeatureSession) CheckVoteByAddress() (common.Address, []*big.Int, error) {
	return _Feature.Contract.CheckVoteByAddress(&_Feature.CallOpts)
}

// CheckVoteByAddress is a free data retrieval call binding the contract method 0xa22b4256.
//
// Solidity: function checkVoteByAddress() view returns(address, uint256[])
func (_Feature *FeatureCallerSession) CheckVoteByAddress() (common.Address, []*big.Int, error) {
	return _Feature.Contract.CheckVoteByAddress(&_Feature.CallOpts)
}

// CountVoteNumber is a free data retrieval call binding the contract method 0xbe8b58c0.
//
// Solidity: function countVoteNumber(uint256 number) view returns(uint256)
func (_Feature *FeatureCaller) CountVoteNumber(opts *bind.CallOpts, number *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Feature.contract.Call(opts, &out, "countVoteNumber", number)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountVoteNumber is a free data retrieval call binding the contract method 0xbe8b58c0.
//
// Solidity: function countVoteNumber(uint256 number) view returns(uint256)
func (_Feature *FeatureSession) CountVoteNumber(number *big.Int) (*big.Int, error) {
	return _Feature.Contract.CountVoteNumber(&_Feature.CallOpts, number)
}

// CountVoteNumber is a free data retrieval call binding the contract method 0xbe8b58c0.
//
// Solidity: function countVoteNumber(uint256 number) view returns(uint256)
func (_Feature *FeatureCallerSession) CountVoteNumber(number *big.Int) (*big.Int, error) {
	return _Feature.Contract.CountVoteNumber(&_Feature.CallOpts, number)
}

// RedeemReward is a paid mutator transaction binding the contract method 0x741f76cd.
//
// Solidity: function redeemReward() returns()
func (_Feature *FeatureTransactor) RedeemReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feature.contract.Transact(opts, "redeemReward")
}

// RedeemReward is a paid mutator transaction binding the contract method 0x741f76cd.
//
// Solidity: function redeemReward() returns()
func (_Feature *FeatureSession) RedeemReward() (*types.Transaction, error) {
	return _Feature.Contract.RedeemReward(&_Feature.TransactOpts)
}

// RedeemReward is a paid mutator transaction binding the contract method 0x741f76cd.
//
// Solidity: function redeemReward() returns()
func (_Feature *FeatureTransactorSession) RedeemReward() (*types.Transaction, error) {
	return _Feature.Contract.RedeemReward(&_Feature.TransactOpts)
}

// SetJackpot is a paid mutator transaction binding the contract method 0x59992cc8.
//
// Solidity: function setJackpot(uint256 number) returns()
func (_Feature *FeatureTransactor) SetJackpot(opts *bind.TransactOpts, number *big.Int) (*types.Transaction, error) {
	return _Feature.contract.Transact(opts, "setJackpot", number)
}

// SetJackpot is a paid mutator transaction binding the contract method 0x59992cc8.
//
// Solidity: function setJackpot(uint256 number) returns()
func (_Feature *FeatureSession) SetJackpot(number *big.Int) (*types.Transaction, error) {
	return _Feature.Contract.SetJackpot(&_Feature.TransactOpts, number)
}

// SetJackpot is a paid mutator transaction binding the contract method 0x59992cc8.
//
// Solidity: function setJackpot(uint256 number) returns()
func (_Feature *FeatureTransactorSession) SetJackpot(number *big.Int) (*types.Transaction, error) {
	return _Feature.Contract.SetJackpot(&_Feature.TransactOpts, number)
}

// VoteTo is a paid mutator transaction binding the contract method 0x3d7b995c.
//
// Solidity: function voteTo(uint256 number) returns()
func (_Feature *FeatureTransactor) VoteTo(opts *bind.TransactOpts, number *big.Int) (*types.Transaction, error) {
	return _Feature.contract.Transact(opts, "voteTo", number)
}

// VoteTo is a paid mutator transaction binding the contract method 0x3d7b995c.
//
// Solidity: function voteTo(uint256 number) returns()
func (_Feature *FeatureSession) VoteTo(number *big.Int) (*types.Transaction, error) {
	return _Feature.Contract.VoteTo(&_Feature.TransactOpts, number)
}

// VoteTo is a paid mutator transaction binding the contract method 0x3d7b995c.
//
// Solidity: function voteTo(uint256 number) returns()
func (_Feature *FeatureTransactorSession) VoteTo(number *big.Int) (*types.Transaction, error) {
	return _Feature.Contract.VoteTo(&_Feature.TransactOpts, number)
}
