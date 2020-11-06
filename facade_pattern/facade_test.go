package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var wf = newWalletFacade("asliddin", 12345)

func TestAccountCheck(t *testing.T) {
	assert.Equal(t, nil, wf.account.checkAccount("asliddin"))
	assert.NotEqual(t, nil, wf.account.checkAccount("unknown"))
}

func TestSecurityCodeCheck(t *testing.T) {
	assert.Equal(t, nil, wf.securityCode.checkCode(12345))
	assert.NotEqual(t, nil, wf.securityCode.checkCode(3902930))
}

func TestSecurityCodeAndAccountCheck(t *testing.T) {
	assert.Equal(t, nil, wf.checkAccountAndCode("asliddin", 12345))
	assert.NotEqual(t, nil, wf.checkAccountAndCode("asliddinbek", 12345))
	assert.NotEqual(t, nil, wf.checkAccountAndCode("asliddin", 392032))
	assert.NotEqual(t, nil, wf.checkAccountAndCode("unkown", 1291))
}

func TestCreditDebitBalance(t *testing.T) {
	assert.Equal(t, nil, wf.addMoneyToWallet("asliddin", 12345, 1000.0))
	assert.Equal(t, nil, wf.deductMoneyFromWallet("asliddin", 12345, 1000.0))
	assert.NotEqual(t, nil, wf.deductMoneyFromWallet("asliddin", 12345, 1000.0))
}
