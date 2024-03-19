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
	got := wal.ConvertToDollars(amt)

	// assert
	assert.Equal(t, want, got)
}

func TestGetBalanceMethod(t *testing.T) {
	amt := 90.00

	wal, _ := NewWallet(amt)

	assert.Equal(t, amt, wal.GetBalance())
}

func TestGetBalanceInUSDMethod(t *testing.T) {
	amt := 82.23
	wal, _ := NewWallet(amt)

	assert.Equal(t, wal.GetBalanceInDollars(), 1.00)
}

func TestShouldVerifyCorrectBalanceAfterDepositINR(t *testing.T) {
	amt := 20.00
	wal, _ := NewWallet(amt)
	dep := 10.00
	want := 30.00

	err := wal.Deposit(dep, "INR")

	got := wal.GetBalance()

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestShoulVerifyDepositMethodThrowsErrorWHenCalledWithInvalidValuesINR(t *testing.T) {
	amt := 25.00
	wal, _ := NewWallet(amt)

	dep := -10.00
	err := wal.Deposit(dep, "INR")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), ErrCannotDepositNegativeAmount)
}

func TestShouldVerifyBalanceOfTheWalletAfterWithdrawalINR(t *testing.T) {
	amt := 25.00
	wal, _ := NewWallet(amt)
	want := 15.00

	with := 10.00

	err := wal.Withdraw(with, "INR")
	got := wal.GetBalance()

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestShouldReturnCorrectBalanceInDollars(t *testing.T) {
	want := 1.00
	wal, _ := NewWallet(82.47)

	got := wal.GetBalanceInDollars()

	assert.Equal(t, want, got)
}

func TestShouldReturnErrorWithWithdrawalAmtGTBalanceINR(t *testing.T) {
	amt := 25.00
	wal, _ := NewWallet(amt)

	with := 30.00
	err := wal.Withdraw(with, "INR")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), ErrInsufficientBalance)
}

func TestShouldReturnErrorWithInvalidValuesForWithdrawalINR(t *testing.T) {
	amt := 25.00
	wal, _ := NewWallet(amt)

	with := -10.00
	err := wal.Withdraw(with, "INR")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), ErrCannotWithdrawNegativeAmount)
}
