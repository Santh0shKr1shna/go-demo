package wallet

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewWalletConstructorMethodWithValidValues(t *testing.T) {
	amt := 98.05

	wal, err := NewWallet(amt)

	assert.Nil(t, err)
	assert.NotNil(t, wal)
}

func TestNewWalletConstructorMethodWithInvalidValues(t *testing.T) {
	amt := -98.05

	wal, err := NewWallet(amt)

	assert.NotNil(t, err)
	assert.Nil(t, wal)
}

func TestConvertToDollarMethod(t *testing.T) {
	amt := 82.47
	want := 1.00

	wal, _ := NewWallet(amt)

	// act
	got := wal.ConvertToDollars()

	// assert
	assert.Equal(t, want, got)
}

func TestGetBalanceMethod(t *testing.T) {
	amt := 90.00

	wal, _ := NewWallet(amt)

	assert.Equal(t, amt, wal.GetBalance())
}

func TestShouldReturnNoErrorWithValidValues(t *testing.T) {
	amt := 25.00
	wal, _ := NewWallet(amt)

	dep := 10.00
	err := wal.Deposit(dep)

	assert.Nil(t, err)
}

func TestShouldReturnErrorWithInvalidValues(t *testing.T) {
	amt := 25.00
	wal, _ := NewWallet(amt)

	dep := -10.00
	err := wal.Deposit(dep)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), ErrCannotDepositNegativeAmount)
}

func TestShouldReturnCorrectBalanceWithValidValues(t *testing.T) {
	amt := 25.00
	wal, _ := NewWallet(amt)
	want := 15.00

	with := 10.00
	err := wal.Withdraw(with)
	got := wal.GetBalance()

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestShouldReturnErrorWithWithdrawalAmtGTBalance(t *testing.T) {
	amt := 25.00
	wal, _ := NewWallet(amt)

	with := 30.00
	err := wal.Withdraw(with)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), ErrInsufficientBalance)
}

func TestShouldReturnErrorWithInvalidValuesForWithDrawal(t *testing.T) {
	amt := 25.00
	wal, _ := NewWallet(amt)

	with := -10.00
	err := wal.Withdraw(with)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), ErrCannotWithdrawNegativeAmount)
}