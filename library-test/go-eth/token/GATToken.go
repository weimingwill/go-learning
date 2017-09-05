// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"resume\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bonus\",\"type\":\"uint256\"}],\"name\":\"setBonus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokensPerKEther\",\"type\":\"uint256\"}],\"name\":\"setTokensPerKEther\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_startTime\",\"type\":\"uint256\"},{\"name\":\"_endTime\",\"type\":\"uint256\"}],\"name\":\"setTimeWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PRESALE_BONUS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_PRIVATE_SALE_CAP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DECIMALS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"END_TIME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contributionMinimum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalTokensSold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserve1Address\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"suspended\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bonus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bankAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_FOUNDATION_CAP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_RESERVE1_CAP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DECIMALSFACTOR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_RESERVE2_CAP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimContractTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserve2Address\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokensPerKEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_FUTURE_CAP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRIBUTION_MIN\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_PUBLIC_SALE_CAP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_TOTAL_CAP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contributionMinimum\",\"type\":\"uint256\"}],\"name\":\"setContributionMinimum\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fundingAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"START_TIME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"suspend\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_PRESALE_CAP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKENS_PER_KETHER\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SYMBOL\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_bankAddress\",\"type\":\"address\"},{\"name\":\"_fundingAddress\",\"type\":\"address\"},{\"name\":\"_reserve1Address\",\"type\":\"address\"},{\"name\":\"_reserve2Address\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"TokensPerKEtherUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"ContributionMinimumUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"BonusAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newStartTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newEndTime\",\"type\":\"uint256\"}],\"name\":\"TimeWindowUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SaleSuspended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SaleResumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"TokenFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ContractTokensReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// CONTRIBUTIONMIN is a free data retrieval call binding the contract method 0xba9bb827.
//
// Solidity: function CONTRIBUTION_MIN() constant returns(uint256)
func (_Token *TokenCaller) CONTRIBUTIONMIN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "CONTRIBUTION_MIN")
	return *ret0, err
}

// CONTRIBUTIONMIN is a free data retrieval call binding the contract method 0xba9bb827.
//
// Solidity: function CONTRIBUTION_MIN() constant returns(uint256)
func (_Token *TokenSession) CONTRIBUTIONMIN() (*big.Int, error) {
	return _Token.Contract.CONTRIBUTIONMIN(&_Token.CallOpts)
}

// CONTRIBUTIONMIN is a free data retrieval call binding the contract method 0xba9bb827.
//
// Solidity: function CONTRIBUTION_MIN() constant returns(uint256)
func (_Token *TokenCallerSession) CONTRIBUTIONMIN() (*big.Int, error) {
	return _Token.Contract.CONTRIBUTIONMIN(&_Token.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() constant returns(uint256)
func (_Token *TokenCaller) DECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "DECIMALS")
	return *ret0, err
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() constant returns(uint256)
func (_Token *TokenSession) DECIMALS() (*big.Int, error) {
	return _Token.Contract.DECIMALS(&_Token.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() constant returns(uint256)
func (_Token *TokenCallerSession) DECIMALS() (*big.Int, error) {
	return _Token.Contract.DECIMALS(&_Token.CallOpts)
}

// DECIMALSFACTOR is a free data retrieval call binding the contract method 0x8bc04eb7.
//
// Solidity: function DECIMALSFACTOR() constant returns(uint256)
func (_Token *TokenCaller) DECIMALSFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "DECIMALSFACTOR")
	return *ret0, err
}

// DECIMALSFACTOR is a free data retrieval call binding the contract method 0x8bc04eb7.
//
// Solidity: function DECIMALSFACTOR() constant returns(uint256)
func (_Token *TokenSession) DECIMALSFACTOR() (*big.Int, error) {
	return _Token.Contract.DECIMALSFACTOR(&_Token.CallOpts)
}

// DECIMALSFACTOR is a free data retrieval call binding the contract method 0x8bc04eb7.
//
// Solidity: function DECIMALSFACTOR() constant returns(uint256)
func (_Token *TokenCallerSession) DECIMALSFACTOR() (*big.Int, error) {
	return _Token.Contract.DECIMALSFACTOR(&_Token.CallOpts)
}

// ENDTIME is a free data retrieval call binding the contract method 0x37ba682d.
//
// Solidity: function END_TIME() constant returns(uint256)
func (_Token *TokenCaller) ENDTIME(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "END_TIME")
	return *ret0, err
}

// ENDTIME is a free data retrieval call binding the contract method 0x37ba682d.
//
// Solidity: function END_TIME() constant returns(uint256)
func (_Token *TokenSession) ENDTIME() (*big.Int, error) {
	return _Token.Contract.ENDTIME(&_Token.CallOpts)
}

// ENDTIME is a free data retrieval call binding the contract method 0x37ba682d.
//
// Solidity: function END_TIME() constant returns(uint256)
func (_Token *TokenCallerSession) ENDTIME() (*big.Int, error) {
	return _Token.Contract.ENDTIME(&_Token.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_Token *TokenCaller) NAME(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "NAME")
	return *ret0, err
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_Token *TokenSession) NAME() (string, error) {
	return _Token.Contract.NAME(&_Token.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_Token *TokenCallerSession) NAME() (string, error) {
	return _Token.Contract.NAME(&_Token.CallOpts)
}

// PRESALEBONUS is a free data retrieval call binding the contract method 0x1b3fddb8.
//
// Solidity: function PRESALE_BONUS() constant returns(uint256)
func (_Token *TokenCaller) PRESALEBONUS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "PRESALE_BONUS")
	return *ret0, err
}

// PRESALEBONUS is a free data retrieval call binding the contract method 0x1b3fddb8.
//
// Solidity: function PRESALE_BONUS() constant returns(uint256)
func (_Token *TokenSession) PRESALEBONUS() (*big.Int, error) {
	return _Token.Contract.PRESALEBONUS(&_Token.CallOpts)
}

// PRESALEBONUS is a free data retrieval call binding the contract method 0x1b3fddb8.
//
// Solidity: function PRESALE_BONUS() constant returns(uint256)
func (_Token *TokenCallerSession) PRESALEBONUS() (*big.Int, error) {
	return _Token.Contract.PRESALEBONUS(&_Token.CallOpts)
}

// STARTTIME is a free data retrieval call binding the contract method 0xddaa26ad.
//
// Solidity: function START_TIME() constant returns(uint256)
func (_Token *TokenCaller) STARTTIME(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "START_TIME")
	return *ret0, err
}

// STARTTIME is a free data retrieval call binding the contract method 0xddaa26ad.
//
// Solidity: function START_TIME() constant returns(uint256)
func (_Token *TokenSession) STARTTIME() (*big.Int, error) {
	return _Token.Contract.STARTTIME(&_Token.CallOpts)
}

// STARTTIME is a free data retrieval call binding the contract method 0xddaa26ad.
//
// Solidity: function START_TIME() constant returns(uint256)
func (_Token *TokenCallerSession) STARTTIME() (*big.Int, error) {
	return _Token.Contract.STARTTIME(&_Token.CallOpts)
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() constant returns(string)
func (_Token *TokenCaller) SYMBOL(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "SYMBOL")
	return *ret0, err
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() constant returns(string)
func (_Token *TokenSession) SYMBOL() (string, error) {
	return _Token.Contract.SYMBOL(&_Token.CallOpts)
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() constant returns(string)
func (_Token *TokenCallerSession) SYMBOL() (string, error) {
	return _Token.Contract.SYMBOL(&_Token.CallOpts)
}

// TOKENSPERKETHER is a free data retrieval call binding the contract method 0xf527c856.
//
// Solidity: function TOKENS_PER_KETHER() constant returns(uint256)
func (_Token *TokenCaller) TOKENSPERKETHER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "TOKENS_PER_KETHER")
	return *ret0, err
}

// TOKENSPERKETHER is a free data retrieval call binding the contract method 0xf527c856.
//
// Solidity: function TOKENS_PER_KETHER() constant returns(uint256)
func (_Token *TokenSession) TOKENSPERKETHER() (*big.Int, error) {
	return _Token.Contract.TOKENSPERKETHER(&_Token.CallOpts)
}

// TOKENSPERKETHER is a free data retrieval call binding the contract method 0xf527c856.
//
// Solidity: function TOKENS_PER_KETHER() constant returns(uint256)
func (_Token *TokenCallerSession) TOKENSPERKETHER() (*big.Int, error) {
	return _Token.Contract.TOKENSPERKETHER(&_Token.CallOpts)
}

// TOKENFOUNDATIONCAP is a free data retrieval call binding the contract method 0x84fd7ef0.
//
// Solidity: function TOKEN_FOUNDATION_CAP() constant returns(uint256)
func (_Token *TokenCaller) TOKENFOUNDATIONCAP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "TOKEN_FOUNDATION_CAP")
	return *ret0, err
}

// TOKENFOUNDATIONCAP is a free data retrieval call binding the contract method 0x84fd7ef0.
//
// Solidity: function TOKEN_FOUNDATION_CAP() constant returns(uint256)
func (_Token *TokenSession) TOKENFOUNDATIONCAP() (*big.Int, error) {
	return _Token.Contract.TOKENFOUNDATIONCAP(&_Token.CallOpts)
}

// TOKENFOUNDATIONCAP is a free data retrieval call binding the contract method 0x84fd7ef0.
//
// Solidity: function TOKEN_FOUNDATION_CAP() constant returns(uint256)
func (_Token *TokenCallerSession) TOKENFOUNDATIONCAP() (*big.Int, error) {
	return _Token.Contract.TOKENFOUNDATIONCAP(&_Token.CallOpts)
}

// TOKENFUTURECAP is a free data retrieval call binding the contract method 0xb46eeebb.
//
// Solidity: function TOKEN_FUTURE_CAP() constant returns(uint256)
func (_Token *TokenCaller) TOKENFUTURECAP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "TOKEN_FUTURE_CAP")
	return *ret0, err
}

// TOKENFUTURECAP is a free data retrieval call binding the contract method 0xb46eeebb.
//
// Solidity: function TOKEN_FUTURE_CAP() constant returns(uint256)
func (_Token *TokenSession) TOKENFUTURECAP() (*big.Int, error) {
	return _Token.Contract.TOKENFUTURECAP(&_Token.CallOpts)
}

// TOKENFUTURECAP is a free data retrieval call binding the contract method 0xb46eeebb.
//
// Solidity: function TOKEN_FUTURE_CAP() constant returns(uint256)
func (_Token *TokenCallerSession) TOKENFUTURECAP() (*big.Int, error) {
	return _Token.Contract.TOKENFUTURECAP(&_Token.CallOpts)
}

// TOKENPRESALECAP is a free data retrieval call binding the contract method 0xeb75dc03.
//
// Solidity: function TOKEN_PRESALE_CAP() constant returns(uint256)
func (_Token *TokenCaller) TOKENPRESALECAP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "TOKEN_PRESALE_CAP")
	return *ret0, err
}

// TOKENPRESALECAP is a free data retrieval call binding the contract method 0xeb75dc03.
//
// Solidity: function TOKEN_PRESALE_CAP() constant returns(uint256)
func (_Token *TokenSession) TOKENPRESALECAP() (*big.Int, error) {
	return _Token.Contract.TOKENPRESALECAP(&_Token.CallOpts)
}

// TOKENPRESALECAP is a free data retrieval call binding the contract method 0xeb75dc03.
//
// Solidity: function TOKEN_PRESALE_CAP() constant returns(uint256)
func (_Token *TokenCallerSession) TOKENPRESALECAP() (*big.Int, error) {
	return _Token.Contract.TOKENPRESALECAP(&_Token.CallOpts)
}

// TOKENPRIVATESALECAP is a free data retrieval call binding the contract method 0x2a9d04f0.
//
// Solidity: function TOKEN_PRIVATE_SALE_CAP() constant returns(uint256)
func (_Token *TokenCaller) TOKENPRIVATESALECAP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "TOKEN_PRIVATE_SALE_CAP")
	return *ret0, err
}

// TOKENPRIVATESALECAP is a free data retrieval call binding the contract method 0x2a9d04f0.
//
// Solidity: function TOKEN_PRIVATE_SALE_CAP() constant returns(uint256)
func (_Token *TokenSession) TOKENPRIVATESALECAP() (*big.Int, error) {
	return _Token.Contract.TOKENPRIVATESALECAP(&_Token.CallOpts)
}

// TOKENPRIVATESALECAP is a free data retrieval call binding the contract method 0x2a9d04f0.
//
// Solidity: function TOKEN_PRIVATE_SALE_CAP() constant returns(uint256)
func (_Token *TokenCallerSession) TOKENPRIVATESALECAP() (*big.Int, error) {
	return _Token.Contract.TOKENPRIVATESALECAP(&_Token.CallOpts)
}

// TOKENPUBLICSALECAP is a free data retrieval call binding the contract method 0xc57a4a4d.
//
// Solidity: function TOKEN_PUBLIC_SALE_CAP() constant returns(uint256)
func (_Token *TokenCaller) TOKENPUBLICSALECAP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "TOKEN_PUBLIC_SALE_CAP")
	return *ret0, err
}

// TOKENPUBLICSALECAP is a free data retrieval call binding the contract method 0xc57a4a4d.
//
// Solidity: function TOKEN_PUBLIC_SALE_CAP() constant returns(uint256)
func (_Token *TokenSession) TOKENPUBLICSALECAP() (*big.Int, error) {
	return _Token.Contract.TOKENPUBLICSALECAP(&_Token.CallOpts)
}

// TOKENPUBLICSALECAP is a free data retrieval call binding the contract method 0xc57a4a4d.
//
// Solidity: function TOKEN_PUBLIC_SALE_CAP() constant returns(uint256)
func (_Token *TokenCallerSession) TOKENPUBLICSALECAP() (*big.Int, error) {
	return _Token.Contract.TOKENPUBLICSALECAP(&_Token.CallOpts)
}

// TOKENRESERVE1CAP is a free data retrieval call binding the contract method 0x8945a8af.
//
// Solidity: function TOKEN_RESERVE1_CAP() constant returns(uint256)
func (_Token *TokenCaller) TOKENRESERVE1CAP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "TOKEN_RESERVE1_CAP")
	return *ret0, err
}

// TOKENRESERVE1CAP is a free data retrieval call binding the contract method 0x8945a8af.
//
// Solidity: function TOKEN_RESERVE1_CAP() constant returns(uint256)
func (_Token *TokenSession) TOKENRESERVE1CAP() (*big.Int, error) {
	return _Token.Contract.TOKENRESERVE1CAP(&_Token.CallOpts)
}

// TOKENRESERVE1CAP is a free data retrieval call binding the contract method 0x8945a8af.
//
// Solidity: function TOKEN_RESERVE1_CAP() constant returns(uint256)
func (_Token *TokenCallerSession) TOKENRESERVE1CAP() (*big.Int, error) {
	return _Token.Contract.TOKENRESERVE1CAP(&_Token.CallOpts)
}

// TOKENRESERVE2CAP is a free data retrieval call binding the contract method 0x8d71f131.
//
// Solidity: function TOKEN_RESERVE2_CAP() constant returns(uint256)
func (_Token *TokenCaller) TOKENRESERVE2CAP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "TOKEN_RESERVE2_CAP")
	return *ret0, err
}

// TOKENRESERVE2CAP is a free data retrieval call binding the contract method 0x8d71f131.
//
// Solidity: function TOKEN_RESERVE2_CAP() constant returns(uint256)
func (_Token *TokenSession) TOKENRESERVE2CAP() (*big.Int, error) {
	return _Token.Contract.TOKENRESERVE2CAP(&_Token.CallOpts)
}

// TOKENRESERVE2CAP is a free data retrieval call binding the contract method 0x8d71f131.
//
// Solidity: function TOKEN_RESERVE2_CAP() constant returns(uint256)
func (_Token *TokenCallerSession) TOKENRESERVE2CAP() (*big.Int, error) {
	return _Token.Contract.TOKENRESERVE2CAP(&_Token.CallOpts)
}

// TOKENTOTALCAP is a free data retrieval call binding the contract method 0xc806a91d.
//
// Solidity: function TOKEN_TOTAL_CAP() constant returns(uint256)
func (_Token *TokenCaller) TOKENTOTALCAP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "TOKEN_TOTAL_CAP")
	return *ret0, err
}

// TOKENTOTALCAP is a free data retrieval call binding the contract method 0xc806a91d.
//
// Solidity: function TOKEN_TOTAL_CAP() constant returns(uint256)
func (_Token *TokenSession) TOKENTOTALCAP() (*big.Int, error) {
	return _Token.Contract.TOKENTOTALCAP(&_Token.CallOpts)
}

// TOKENTOTALCAP is a free data retrieval call binding the contract method 0xc806a91d.
//
// Solidity: function TOKEN_TOTAL_CAP() constant returns(uint256)
func (_Token *TokenCallerSession) TOKENTOTALCAP() (*big.Int, error) {
	return _Token.Contract.TOKENTOTALCAP(&_Token.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_Token *TokenCaller) BankAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "bankAddress")
	return *ret0, err
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_Token *TokenSession) BankAddress() (common.Address, error) {
	return _Token.Contract.BankAddress(&_Token.CallOpts)
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_Token *TokenCallerSession) BankAddress() (common.Address, error) {
	return _Token.Contract.BankAddress(&_Token.CallOpts)
}

// Bonus is a free data retrieval call binding the contract method 0x75b4d78c.
//
// Solidity: function bonus() constant returns(uint256)
func (_Token *TokenCaller) Bonus(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "bonus")
	return *ret0, err
}

// Bonus is a free data retrieval call binding the contract method 0x75b4d78c.
//
// Solidity: function bonus() constant returns(uint256)
func (_Token *TokenSession) Bonus() (*big.Int, error) {
	return _Token.Contract.Bonus(&_Token.CallOpts)
}

// Bonus is a free data retrieval call binding the contract method 0x75b4d78c.
//
// Solidity: function bonus() constant returns(uint256)
func (_Token *TokenCallerSession) Bonus() (*big.Int, error) {
	return _Token.Contract.Bonus(&_Token.CallOpts)
}

// ContributionMinimum is a free data retrieval call binding the contract method 0x4f2d7ab5.
//
// Solidity: function contributionMinimum() constant returns(uint256)
func (_Token *TokenCaller) ContributionMinimum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "contributionMinimum")
	return *ret0, err
}

// ContributionMinimum is a free data retrieval call binding the contract method 0x4f2d7ab5.
//
// Solidity: function contributionMinimum() constant returns(uint256)
func (_Token *TokenSession) ContributionMinimum() (*big.Int, error) {
	return _Token.Contract.ContributionMinimum(&_Token.CallOpts)
}

// ContributionMinimum is a free data retrieval call binding the contract method 0x4f2d7ab5.
//
// Solidity: function contributionMinimum() constant returns(uint256)
func (_Token *TokenCallerSession) ContributionMinimum() (*big.Int, error) {
	return _Token.Contract.ContributionMinimum(&_Token.CallOpts)
}

// CurrentTime is a free data retrieval call binding the contract method 0xd18e81b3.
//
// Solidity: function currentTime() constant returns(uint256)
func (_Token *TokenCaller) CurrentTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "currentTime")
	return *ret0, err
}

// CurrentTime is a free data retrieval call binding the contract method 0xd18e81b3.
//
// Solidity: function currentTime() constant returns(uint256)
func (_Token *TokenSession) CurrentTime() (*big.Int, error) {
	return _Token.Contract.CurrentTime(&_Token.CallOpts)
}

// CurrentTime is a free data retrieval call binding the contract method 0xd18e81b3.
//
// Solidity: function currentTime() constant returns(uint256)
func (_Token *TokenCallerSession) CurrentTime() (*big.Int, error) {
	return _Token.Contract.CurrentTime(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Token *TokenSession) Decimals() (*big.Int, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Token *TokenCallerSession) Decimals() (*big.Int, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_Token *TokenCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "endTime")
	return *ret0, err
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_Token *TokenSession) EndTime() (*big.Int, error) {
	return _Token.Contract.EndTime(&_Token.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_Token *TokenCallerSession) EndTime() (*big.Int, error) {
	return _Token.Contract.EndTime(&_Token.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() constant returns(bool)
func (_Token *TokenCaller) Finalized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "finalized")
	return *ret0, err
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() constant returns(bool)
func (_Token *TokenSession) Finalized() (bool, error) {
	return _Token.Contract.Finalized(&_Token.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() constant returns(bool)
func (_Token *TokenCallerSession) Finalized() (bool, error) {
	return _Token.Contract.Finalized(&_Token.CallOpts)
}

// FundingAddress is a free data retrieval call binding the contract method 0xd3b7bfb4.
//
// Solidity: function fundingAddress() constant returns(address)
func (_Token *TokenCaller) FundingAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "fundingAddress")
	return *ret0, err
}

// FundingAddress is a free data retrieval call binding the contract method 0xd3b7bfb4.
//
// Solidity: function fundingAddress() constant returns(address)
func (_Token *TokenSession) FundingAddress() (common.Address, error) {
	return _Token.Contract.FundingAddress(&_Token.CallOpts)
}

// FundingAddress is a free data retrieval call binding the contract method 0xd3b7bfb4.
//
// Solidity: function fundingAddress() constant returns(address)
func (_Token *TokenCallerSession) FundingAddress() (common.Address, error) {
	return _Token.Contract.FundingAddress(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Token *TokenCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "newOwner")
	return *ret0, err
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Token *TokenSession) NewOwner() (common.Address, error) {
	return _Token.Contract.NewOwner(&_Token.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Token *TokenCallerSession) NewOwner() (common.Address, error) {
	return _Token.Contract.NewOwner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Reserve1Address is a free data retrieval call binding the contract method 0x63cb2afb.
//
// Solidity: function reserve1Address() constant returns(address)
func (_Token *TokenCaller) Reserve1Address(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "reserve1Address")
	return *ret0, err
}

// Reserve1Address is a free data retrieval call binding the contract method 0x63cb2afb.
//
// Solidity: function reserve1Address() constant returns(address)
func (_Token *TokenSession) Reserve1Address() (common.Address, error) {
	return _Token.Contract.Reserve1Address(&_Token.CallOpts)
}

// Reserve1Address is a free data retrieval call binding the contract method 0x63cb2afb.
//
// Solidity: function reserve1Address() constant returns(address)
func (_Token *TokenCallerSession) Reserve1Address() (common.Address, error) {
	return _Token.Contract.Reserve1Address(&_Token.CallOpts)
}

// Reserve2Address is a free data retrieval call binding the contract method 0x979260bd.
//
// Solidity: function reserve2Address() constant returns(address)
func (_Token *TokenCaller) Reserve2Address(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "reserve2Address")
	return *ret0, err
}

// Reserve2Address is a free data retrieval call binding the contract method 0x979260bd.
//
// Solidity: function reserve2Address() constant returns(address)
func (_Token *TokenSession) Reserve2Address() (common.Address, error) {
	return _Token.Contract.Reserve2Address(&_Token.CallOpts)
}

// Reserve2Address is a free data retrieval call binding the contract method 0x979260bd.
//
// Solidity: function reserve2Address() constant returns(address)
func (_Token *TokenCallerSession) Reserve2Address() (common.Address, error) {
	return _Token.Contract.Reserve2Address(&_Token.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_Token *TokenCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "startTime")
	return *ret0, err
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_Token *TokenSession) StartTime() (*big.Int, error) {
	return _Token.Contract.StartTime(&_Token.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_Token *TokenCallerSession) StartTime() (*big.Int, error) {
	return _Token.Contract.StartTime(&_Token.CallOpts)
}

// Suspended is a free data retrieval call binding the contract method 0x702efdf3.
//
// Solidity: function suspended() constant returns(bool)
func (_Token *TokenCaller) Suspended(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "suspended")
	return *ret0, err
}

// Suspended is a free data retrieval call binding the contract method 0x702efdf3.
//
// Solidity: function suspended() constant returns(bool)
func (_Token *TokenSession) Suspended() (bool, error) {
	return _Token.Contract.Suspended(&_Token.CallOpts)
}

// Suspended is a free data retrieval call binding the contract method 0x702efdf3.
//
// Solidity: function suspended() constant returns(bool)
func (_Token *TokenCallerSession) Suspended() (bool, error) {
	return _Token.Contract.Suspended(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TokensPerKEther is a free data retrieval call binding the contract method 0xa5bc770c.
//
// Solidity: function tokensPerKEther() constant returns(uint256)
func (_Token *TokenCaller) TokensPerKEther(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "tokensPerKEther")
	return *ret0, err
}

// TokensPerKEther is a free data retrieval call binding the contract method 0xa5bc770c.
//
// Solidity: function tokensPerKEther() constant returns(uint256)
func (_Token *TokenSession) TokensPerKEther() (*big.Int, error) {
	return _Token.Contract.TokensPerKEther(&_Token.CallOpts)
}

// TokensPerKEther is a free data retrieval call binding the contract method 0xa5bc770c.
//
// Solidity: function tokensPerKEther() constant returns(uint256)
func (_Token *TokenCallerSession) TokensPerKEther() (*big.Int, error) {
	return _Token.Contract.TokensPerKEther(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalTokensSold is a free data retrieval call binding the contract method 0x63b20117.
//
// Solidity: function totalTokensSold() constant returns(uint256)
func (_Token *TokenCaller) TotalTokensSold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalTokensSold")
	return *ret0, err
}

// TotalTokensSold is a free data retrieval call binding the contract method 0x63b20117.
//
// Solidity: function totalTokensSold() constant returns(uint256)
func (_Token *TokenSession) TotalTokensSold() (*big.Int, error) {
	return _Token.Contract.TotalTokensSold(&_Token.CallOpts)
}

// TotalTokensSold is a free data retrieval call binding the contract method 0x63b20117.
//
// Solidity: function totalTokensSold() constant returns(uint256)
func (_Token *TokenCallerSession) TotalTokensSold() (*big.Int, error) {
	return _Token.Contract.TotalTokensSold(&_Token.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns(bool)
func (_Token *TokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns(bool)
func (_Token *TokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _Token.Contract.AcceptOwnership(&_Token.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns(bool)
func (_Token *TokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Token.Contract.AcceptOwnership(&_Token.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns(uint256)
func (_Token *TokenTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns(uint256)
func (_Token *TokenSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _Token.Contract.BuyTokens(&_Token.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns(uint256)
func (_Token *TokenTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _Token.Contract.BuyTokens(&_Token.TransactOpts, beneficiary)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_Token *TokenTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_Token *TokenSession) Finalize() (*types.Transaction, error) {
	return _Token.Contract.Finalize(&_Token.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_Token *TokenTransactorSession) Finalize() (*types.Transaction, error) {
	return _Token.Contract.Finalize(&_Token.TransactOpts)
}

// ReclaimContractTokens is a paid mutator transaction binding the contract method 0x8f14d8a3.
//
// Solidity: function reclaimContractTokens() returns(bool)
func (_Token *TokenTransactor) ReclaimContractTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "reclaimContractTokens")
}

// ReclaimContractTokens is a paid mutator transaction binding the contract method 0x8f14d8a3.
//
// Solidity: function reclaimContractTokens() returns(bool)
func (_Token *TokenSession) ReclaimContractTokens() (*types.Transaction, error) {
	return _Token.Contract.ReclaimContractTokens(&_Token.TransactOpts)
}

// ReclaimContractTokens is a paid mutator transaction binding the contract method 0x8f14d8a3.
//
// Solidity: function reclaimContractTokens() returns(bool)
func (_Token *TokenTransactorSession) ReclaimContractTokens() (*types.Transaction, error) {
	return _Token.Contract.ReclaimContractTokens(&_Token.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns(bool)
func (_Token *TokenTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns(bool)
func (_Token *TokenSession) Resume() (*types.Transaction, error) {
	return _Token.Contract.Resume(&_Token.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns(bool)
func (_Token *TokenTransactorSession) Resume() (*types.Transaction, error) {
	return _Token.Contract.Resume(&_Token.TransactOpts)
}

// SetBonus is a paid mutator transaction binding the contract method 0x0b98f975.
//
// Solidity: function setBonus(_bonus uint256) returns(bool)
func (_Token *TokenTransactor) SetBonus(opts *bind.TransactOpts, _bonus *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setBonus", _bonus)
}

// SetBonus is a paid mutator transaction binding the contract method 0x0b98f975.
//
// Solidity: function setBonus(_bonus uint256) returns(bool)
func (_Token *TokenSession) SetBonus(_bonus *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetBonus(&_Token.TransactOpts, _bonus)
}

// SetBonus is a paid mutator transaction binding the contract method 0x0b98f975.
//
// Solidity: function setBonus(_bonus uint256) returns(bool)
func (_Token *TokenTransactorSession) SetBonus(_bonus *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetBonus(&_Token.TransactOpts, _bonus)
}

// SetContributionMinimum is a paid mutator transaction binding the contract method 0xcb0b7b03.
//
// Solidity: function setContributionMinimum(_contributionMinimum uint256) returns(bool)
func (_Token *TokenTransactor) SetContributionMinimum(opts *bind.TransactOpts, _contributionMinimum *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setContributionMinimum", _contributionMinimum)
}

// SetContributionMinimum is a paid mutator transaction binding the contract method 0xcb0b7b03.
//
// Solidity: function setContributionMinimum(_contributionMinimum uint256) returns(bool)
func (_Token *TokenSession) SetContributionMinimum(_contributionMinimum *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetContributionMinimum(&_Token.TransactOpts, _contributionMinimum)
}

// SetContributionMinimum is a paid mutator transaction binding the contract method 0xcb0b7b03.
//
// Solidity: function setContributionMinimum(_contributionMinimum uint256) returns(bool)
func (_Token *TokenTransactorSession) SetContributionMinimum(_contributionMinimum *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetContributionMinimum(&_Token.TransactOpts, _contributionMinimum)
}

// SetTimeWindow is a paid mutator transaction binding the contract method 0x1a79c5de.
//
// Solidity: function setTimeWindow(_startTime uint256, _endTime uint256) returns(bool)
func (_Token *TokenTransactor) SetTimeWindow(opts *bind.TransactOpts, _startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setTimeWindow", _startTime, _endTime)
}

// SetTimeWindow is a paid mutator transaction binding the contract method 0x1a79c5de.
//
// Solidity: function setTimeWindow(_startTime uint256, _endTime uint256) returns(bool)
func (_Token *TokenSession) SetTimeWindow(_startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetTimeWindow(&_Token.TransactOpts, _startTime, _endTime)
}

// SetTimeWindow is a paid mutator transaction binding the contract method 0x1a79c5de.
//
// Solidity: function setTimeWindow(_startTime uint256, _endTime uint256) returns(bool)
func (_Token *TokenTransactorSession) SetTimeWindow(_startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetTimeWindow(&_Token.TransactOpts, _startTime, _endTime)
}

// SetTokensPerKEther is a paid mutator transaction binding the contract method 0x0e9d02cc.
//
// Solidity: function setTokensPerKEther(_tokensPerKEther uint256) returns(bool)
func (_Token *TokenTransactor) SetTokensPerKEther(opts *bind.TransactOpts, _tokensPerKEther *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setTokensPerKEther", _tokensPerKEther)
}

// SetTokensPerKEther is a paid mutator transaction binding the contract method 0x0e9d02cc.
//
// Solidity: function setTokensPerKEther(_tokensPerKEther uint256) returns(bool)
func (_Token *TokenSession) SetTokensPerKEther(_tokensPerKEther *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetTokensPerKEther(&_Token.TransactOpts, _tokensPerKEther)
}

// SetTokensPerKEther is a paid mutator transaction binding the contract method 0x0e9d02cc.
//
// Solidity: function setTokensPerKEther(_tokensPerKEther uint256) returns(bool)
func (_Token *TokenTransactorSession) SetTokensPerKEther(_tokensPerKEther *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetTokensPerKEther(&_Token.TransactOpts, _tokensPerKEther)
}

// Suspend is a paid mutator transaction binding the contract method 0xe6400bbe.
//
// Solidity: function suspend() returns(bool)
func (_Token *TokenTransactor) Suspend(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "suspend")
}

// Suspend is a paid mutator transaction binding the contract method 0xe6400bbe.
//
// Solidity: function suspend() returns(bool)
func (_Token *TokenSession) Suspend() (*types.Transaction, error) {
	return _Token.Contract.Suspend(&_Token.TransactOpts)
}

// Suspend is a paid mutator transaction binding the contract method 0xe6400bbe.
//
// Solidity: function suspend() returns(bool)
func (_Token *TokenTransactorSession) Suspend() (*types.Transaction, error) {
	return _Token.Contract.Suspend(&_Token.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _amount uint256) returns(success bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _amount uint256) returns(success bool)
func (_Token *TokenSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _amount uint256) returns(success bool)
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _amount uint256) returns(success bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _amount uint256) returns(success bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _amount uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Token *TokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Token *TokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Token *TokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, _newOwner)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*TokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// TokenBonusAmountUpdatedIterator is returned from FilterBonusAmountUpdated and is used to iterate over the raw logs and unpacked data for BonusAmountUpdated events raised by the Token contract.
type TokenBonusAmountUpdatedIterator struct {
	Event *TokenBonusAmountUpdated // Event containing the contract specifics and raw log

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
func (it *TokenBonusAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBonusAmountUpdated)
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
		it.Event = new(TokenBonusAmountUpdated)
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
func (it *TokenBonusAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBonusAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBonusAmountUpdated represents a BonusAmountUpdated event raised by the Token contract.
type TokenBonusAmountUpdated struct {
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBonusAmountUpdated is a free log retrieval operation binding the contract event 0x7d9e5243a26ab840171b7338448bae49afbea12ab5209c779f4e2ae6e2d141cd.
//
// Solidity: e BonusAmountUpdated(newAmount uint256)
func (_Token *TokenFilterer) FilterBonusAmountUpdated(opts *bind.FilterOpts) (*TokenBonusAmountUpdatedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "BonusAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenBonusAmountUpdatedIterator{contract: _Token.contract, event: "BonusAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchBonusAmountUpdated is a free log subscription operation binding the contract event 0x7d9e5243a26ab840171b7338448bae49afbea12ab5209c779f4e2ae6e2d141cd.
//
// Solidity: e BonusAmountUpdated(newAmount uint256)
func (_Token *TokenFilterer) WatchBonusAmountUpdated(opts *bind.WatchOpts, sink chan<- *TokenBonusAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "BonusAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBonusAmountUpdated)
				if err := _Token.contract.UnpackLog(event, "BonusAmountUpdated", log); err != nil {
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

// TokenContractTokensReclaimedIterator is returned from FilterContractTokensReclaimed and is used to iterate over the raw logs and unpacked data for ContractTokensReclaimed events raised by the Token contract.
type TokenContractTokensReclaimedIterator struct {
	Event *TokenContractTokensReclaimed // Event containing the contract specifics and raw log

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
func (it *TokenContractTokensReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenContractTokensReclaimed)
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
		it.Event = new(TokenContractTokensReclaimed)
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
func (it *TokenContractTokensReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenContractTokensReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenContractTokensReclaimed represents a ContractTokensReclaimed event raised by the Token contract.
type TokenContractTokensReclaimed struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterContractTokensReclaimed is a free log retrieval operation binding the contract event 0x2bdbc0ce7fbf2aef4c647c03c4bfd8944d985741800d90ca4f1e8c6f5b77419e.
//
// Solidity: e ContractTokensReclaimed(amount uint256)
func (_Token *TokenFilterer) FilterContractTokensReclaimed(opts *bind.FilterOpts) (*TokenContractTokensReclaimedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "ContractTokensReclaimed")
	if err != nil {
		return nil, err
	}
	return &TokenContractTokensReclaimedIterator{contract: _Token.contract, event: "ContractTokensReclaimed", logs: logs, sub: sub}, nil
}

// WatchContractTokensReclaimed is a free log subscription operation binding the contract event 0x2bdbc0ce7fbf2aef4c647c03c4bfd8944d985741800d90ca4f1e8c6f5b77419e.
//
// Solidity: e ContractTokensReclaimed(amount uint256)
func (_Token *TokenFilterer) WatchContractTokensReclaimed(opts *bind.WatchOpts, sink chan<- *TokenContractTokensReclaimed) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "ContractTokensReclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenContractTokensReclaimed)
				if err := _Token.contract.UnpackLog(event, "ContractTokensReclaimed", log); err != nil {
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

// TokenContributionMinimumUpdatedIterator is returned from FilterContributionMinimumUpdated and is used to iterate over the raw logs and unpacked data for ContributionMinimumUpdated events raised by the Token contract.
type TokenContributionMinimumUpdatedIterator struct {
	Event *TokenContributionMinimumUpdated // Event containing the contract specifics and raw log

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
func (it *TokenContributionMinimumUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenContributionMinimumUpdated)
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
		it.Event = new(TokenContributionMinimumUpdated)
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
func (it *TokenContributionMinimumUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenContributionMinimumUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenContributionMinimumUpdated represents a ContributionMinimumUpdated event raised by the Token contract.
type TokenContributionMinimumUpdated struct {
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContributionMinimumUpdated is a free log retrieval operation binding the contract event 0xe6fafef0739724aaa3f73724864d5821481aa094d2c77c7378b77a69e34d9ac7.
//
// Solidity: e ContributionMinimumUpdated(newAmount uint256)
func (_Token *TokenFilterer) FilterContributionMinimumUpdated(opts *bind.FilterOpts) (*TokenContributionMinimumUpdatedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "ContributionMinimumUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenContributionMinimumUpdatedIterator{contract: _Token.contract, event: "ContributionMinimumUpdated", logs: logs, sub: sub}, nil
}

// WatchContributionMinimumUpdated is a free log subscription operation binding the contract event 0xe6fafef0739724aaa3f73724864d5821481aa094d2c77c7378b77a69e34d9ac7.
//
// Solidity: e ContributionMinimumUpdated(newAmount uint256)
func (_Token *TokenFilterer) WatchContributionMinimumUpdated(opts *bind.WatchOpts, sink chan<- *TokenContributionMinimumUpdated) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "ContributionMinimumUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenContributionMinimumUpdated)
				if err := _Token.contract.UnpackLog(event, "ContributionMinimumUpdated", log); err != nil {
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

// TokenOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Token contract.
type TokenOwnerChangedIterator struct {
	Event *TokenOwnerChanged // Event containing the contract specifics and raw log

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
func (it *TokenOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnerChanged)
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
		it.Event = new(TokenOwnerChanged)
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
func (it *TokenOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnerChanged represents a OwnerChanged event raised by the Token contract.
type TokenOwnerChanged struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: e OwnerChanged(_newOwner indexed address)
func (_Token *TokenFilterer) FilterOwnerChanged(opts *bind.FilterOpts, _newOwner []common.Address) (*TokenOwnerChangedIterator, error) {

	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnerChanged", _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenOwnerChangedIterator{contract: _Token.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: e OwnerChanged(_newOwner indexed address)
func (_Token *TokenFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *TokenOwnerChanged, _newOwner []common.Address) (event.Subscription, error) {

	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnerChanged", _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnerChanged)
				if err := _Token.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// TokenSaleResumedIterator is returned from FilterSaleResumed and is used to iterate over the raw logs and unpacked data for SaleResumed events raised by the Token contract.
type TokenSaleResumedIterator struct {
	Event *TokenSaleResumed // Event containing the contract specifics and raw log

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
func (it *TokenSaleResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSaleResumed)
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
		it.Event = new(TokenSaleResumed)
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
func (it *TokenSaleResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSaleResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSaleResumed represents a SaleResumed event raised by the Token contract.
type TokenSaleResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSaleResumed is a free log retrieval operation binding the contract event 0xbcbdbf400d5c713d9679ffa947f717848591ab5a7d1608c49119db603c4942cb.
//
// Solidity: e SaleResumed()
func (_Token *TokenFilterer) FilterSaleResumed(opts *bind.FilterOpts) (*TokenSaleResumedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "SaleResumed")
	if err != nil {
		return nil, err
	}
	return &TokenSaleResumedIterator{contract: _Token.contract, event: "SaleResumed", logs: logs, sub: sub}, nil
}

// WatchSaleResumed is a free log subscription operation binding the contract event 0xbcbdbf400d5c713d9679ffa947f717848591ab5a7d1608c49119db603c4942cb.
//
// Solidity: e SaleResumed()
func (_Token *TokenFilterer) WatchSaleResumed(opts *bind.WatchOpts, sink chan<- *TokenSaleResumed) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "SaleResumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSaleResumed)
				if err := _Token.contract.UnpackLog(event, "SaleResumed", log); err != nil {
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

// TokenSaleSuspendedIterator is returned from FilterSaleSuspended and is used to iterate over the raw logs and unpacked data for SaleSuspended events raised by the Token contract.
type TokenSaleSuspendedIterator struct {
	Event *TokenSaleSuspended // Event containing the contract specifics and raw log

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
func (it *TokenSaleSuspendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSaleSuspended)
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
		it.Event = new(TokenSaleSuspended)
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
func (it *TokenSaleSuspendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSaleSuspendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSaleSuspended represents a SaleSuspended event raised by the Token contract.
type TokenSaleSuspended struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSaleSuspended is a free log retrieval operation binding the contract event 0xe14916b4c867f32e91547d295f9b845b805d5b8c813daa3adbc1597f80a0c5eb.
//
// Solidity: e SaleSuspended()
func (_Token *TokenFilterer) FilterSaleSuspended(opts *bind.FilterOpts) (*TokenSaleSuspendedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "SaleSuspended")
	if err != nil {
		return nil, err
	}
	return &TokenSaleSuspendedIterator{contract: _Token.contract, event: "SaleSuspended", logs: logs, sub: sub}, nil
}

// WatchSaleSuspended is a free log subscription operation binding the contract event 0xe14916b4c867f32e91547d295f9b845b805d5b8c813daa3adbc1597f80a0c5eb.
//
// Solidity: e SaleSuspended()
func (_Token *TokenFilterer) WatchSaleSuspended(opts *bind.WatchOpts, sink chan<- *TokenSaleSuspended) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "SaleSuspended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSaleSuspended)
				if err := _Token.contract.UnpackLog(event, "SaleSuspended", log); err != nil {
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

// TokenTimeWindowUpdatedIterator is returned from FilterTimeWindowUpdated and is used to iterate over the raw logs and unpacked data for TimeWindowUpdated events raised by the Token contract.
type TokenTimeWindowUpdatedIterator struct {
	Event *TokenTimeWindowUpdated // Event containing the contract specifics and raw log

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
func (it *TokenTimeWindowUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTimeWindowUpdated)
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
		it.Event = new(TokenTimeWindowUpdated)
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
func (it *TokenTimeWindowUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTimeWindowUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTimeWindowUpdated represents a TimeWindowUpdated event raised by the Token contract.
type TokenTimeWindowUpdated struct {
	NewStartTime *big.Int
	NewEndTime   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTimeWindowUpdated is a free log retrieval operation binding the contract event 0x6c118f466f3e47773b4c9da27f548aafdf212f592e28574f28ecc67ef19cd451.
//
// Solidity: e TimeWindowUpdated(newStartTime uint256, newEndTime uint256)
func (_Token *TokenFilterer) FilterTimeWindowUpdated(opts *bind.FilterOpts) (*TokenTimeWindowUpdatedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "TimeWindowUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenTimeWindowUpdatedIterator{contract: _Token.contract, event: "TimeWindowUpdated", logs: logs, sub: sub}, nil
}

// WatchTimeWindowUpdated is a free log subscription operation binding the contract event 0x6c118f466f3e47773b4c9da27f548aafdf212f592e28574f28ecc67ef19cd451.
//
// Solidity: e TimeWindowUpdated(newStartTime uint256, newEndTime uint256)
func (_Token *TokenFilterer) WatchTimeWindowUpdated(opts *bind.WatchOpts, sink chan<- *TokenTimeWindowUpdated) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "TimeWindowUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTimeWindowUpdated)
				if err := _Token.contract.UnpackLog(event, "TimeWindowUpdated", log); err != nil {
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

// TokenTokenFinalizedIterator is returned from FilterTokenFinalized and is used to iterate over the raw logs and unpacked data for TokenFinalized events raised by the Token contract.
type TokenTokenFinalizedIterator struct {
	Event *TokenTokenFinalized // Event containing the contract specifics and raw log

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
func (it *TokenTokenFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTokenFinalized)
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
		it.Event = new(TokenTokenFinalized)
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
func (it *TokenTokenFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTokenFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTokenFinalized represents a TokenFinalized event raised by the Token contract.
type TokenTokenFinalized struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTokenFinalized is a free log retrieval operation binding the contract event 0x0f9b481a37d4503bc76152eef0e2ba08850a8db76068c93d4d6bec0395aee723.
//
// Solidity: e TokenFinalized()
func (_Token *TokenFilterer) FilterTokenFinalized(opts *bind.FilterOpts) (*TokenTokenFinalizedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "TokenFinalized")
	if err != nil {
		return nil, err
	}
	return &TokenTokenFinalizedIterator{contract: _Token.contract, event: "TokenFinalized", logs: logs, sub: sub}, nil
}

// WatchTokenFinalized is a free log subscription operation binding the contract event 0x0f9b481a37d4503bc76152eef0e2ba08850a8db76068c93d4d6bec0395aee723.
//
// Solidity: e TokenFinalized()
func (_Token *TokenFilterer) WatchTokenFinalized(opts *bind.WatchOpts, sink chan<- *TokenTokenFinalized) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "TokenFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTokenFinalized)
				if err := _Token.contract.UnpackLog(event, "TokenFinalized", log); err != nil {
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

// TokenTokensPerKEtherUpdatedIterator is returned from FilterTokensPerKEtherUpdated and is used to iterate over the raw logs and unpacked data for TokensPerKEtherUpdated events raised by the Token contract.
type TokenTokensPerKEtherUpdatedIterator struct {
	Event *TokenTokensPerKEtherUpdated // Event containing the contract specifics and raw log

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
func (it *TokenTokensPerKEtherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTokensPerKEtherUpdated)
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
		it.Event = new(TokenTokensPerKEtherUpdated)
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
func (it *TokenTokensPerKEtherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTokensPerKEtherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTokensPerKEtherUpdated represents a TokensPerKEtherUpdated event raised by the Token contract.
type TokenTokensPerKEtherUpdated struct {
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensPerKEtherUpdated is a free log retrieval operation binding the contract event 0xee386bebbe46d39825c2b93313aa1ab1dc57d4774cac81c6debb8c611c9227ab.
//
// Solidity: e TokensPerKEtherUpdated(newAmount uint256)
func (_Token *TokenFilterer) FilterTokensPerKEtherUpdated(opts *bind.FilterOpts) (*TokenTokensPerKEtherUpdatedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "TokensPerKEtherUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenTokensPerKEtherUpdatedIterator{contract: _Token.contract, event: "TokensPerKEtherUpdated", logs: logs, sub: sub}, nil
}

// WatchTokensPerKEtherUpdated is a free log subscription operation binding the contract event 0xee386bebbe46d39825c2b93313aa1ab1dc57d4774cac81c6debb8c611c9227ab.
//
// Solidity: e TokensPerKEtherUpdated(newAmount uint256)
func (_Token *TokenFilterer) WatchTokensPerKEtherUpdated(opts *bind.WatchOpts, sink chan<- *TokenTokensPerKEtherUpdated) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "TokensPerKEtherUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTokensPerKEtherUpdated)
				if err := _Token.contract.UnpackLog(event, "TokensPerKEtherUpdated", log); err != nil {
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

// TokenTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the Token contract.
type TokenTokensPurchasedIterator struct {
	Event *TokenTokensPurchased // Event containing the contract specifics and raw log

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
func (it *TokenTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTokensPurchased)
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
		it.Event = new(TokenTokensPurchased)
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
func (it *TokenTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTokensPurchased represents a TokensPurchased event raised by the Token contract.
type TokenTokensPurchased struct {
	Beneficiary common.Address
	Cost        *big.Int
	Tokens      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x8fafebcaf9d154343dad25669bfa277f4fbacd7ac6b0c4fed522580e040a0f33.
//
// Solidity: e TokensPurchased(beneficiary indexed address, cost uint256, tokens uint256)
func (_Token *TokenFilterer) FilterTokensPurchased(opts *bind.FilterOpts, beneficiary []common.Address) (*TokenTokensPurchasedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "TokensPurchased", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &TokenTokensPurchasedIterator{contract: _Token.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x8fafebcaf9d154343dad25669bfa277f4fbacd7ac6b0c4fed522580e040a0f33.
//
// Solidity: e TokensPurchased(beneficiary indexed address, cost uint256, tokens uint256)
func (_Token *TokenFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *TokenTokensPurchased, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "TokensPurchased", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTokensPurchased)
				if err := _Token.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*TokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
