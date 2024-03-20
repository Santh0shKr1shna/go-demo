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

// TODO: test money constructor functions

func TestMoneyconstructorWithValidValues(t *testing.T) {
	amt := 98.05
	cur := "INR"

	money, err := NewMoney(amt, cur)

	assert.Nil(t, err)
	assert.NotNil(t, money)
}

func TestMoneyconstructorWithNegativeValues(t *testing.T) {
	amt := -98.05
	cur := "INR"

	money, err := NewMoney(amt, cur)

	assert.NotNil(t, err)
	assert.Nil(t, money)
}

func TestMoneyconstructorWithNonSupportedCurrency(t *testing.T) {
	amt := 98.05
	cur := "YEN"

	money, err := NewMoney(amt, cur)

	assert.NotNil(t, err)
	assert.Nil(t, money)
}

func TestConvertRupeeToCurrencyMethodWithnValidValues(t *testing.T) {
	amt := 82.47
	requiredCur := "USD"
	nativeCur := "INR"
	want := 1.00

	wal, _ := NewWallet(5.00)
	money, _ := NewMoney(amt, nativeCur)

	// act
	got, err := wal.ConvertRupeeToCurrency(*money, requiredCur)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestConvertRupeeToCurrencyMethodWithNativeCurrencyNotRupee(t *testing.T) {
	amt := 82.47
	requiredCur := "USD"
	nativeCur := "EUR"

	wal, _ := NewWallet(0.01)
	money, _ := NewMoney(amt, nativeCur)

	// act
	got, err := wal.ConvertRupeeToCurrency(*money, requiredCur)

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, got)
}

func TestShouldVerifyConvertToRupeeMethod(t *testing.T) {
	amt := 1.00
	nativeCur := "EUR"
	want := 90.05

	wal, _ := NewWallet(5.01)
	money, _ := NewMoney(amt, nativeCur)

	// act
	got := wal.ConvertToRupee(*money)

	assert.Equal(t, got, want)
}

func TestShouldVerifyBalanceAfterDeposit(t *testing.T) {
	amt := 25.00
	wal, _ := NewWallet(amt)

	with := 10.00
	withMoney, _ := NewMoney(with, "INR")
	want := 35.00
	
	// Act
	wal.Deposit(*withMoney)
	got := wal.GetBalanceInRupees()

	assert.Equal(t, got, want)
}
func TestGetBalnceAfterMultipleDepositsInRupees(t *testing.T) {
	wal, _ := NewWallet(0.00)
	currency := "INR"

	deposit1, _ := NewMoney(5.00, currency)
	deposit2, _ := NewMoney(15.00, currency)
	deposit3, _ := NewMoney(20.00, currency)

	wal.Deposit(*deposit1)
	wal.Deposit(*deposit2)
	wal.Deposit(*deposit3)

	want := 40.00
	got := wal.GetBalanceInRupees()

	assert.Equal(t, got, want)
}

func TestGetBalnceAfterMultipleDepositsInVariousCurrencies(t *testing.T) {
	wal, _ := NewWallet(0.00)

	deposit1, _ := NewMoney(5.00, "USD")
	deposit2, _ := NewMoney(15.00, "EUR")
	deposit3, _ := NewMoney(20.00, "INR")

	wal.Deposit(*deposit1)
	wal.Deposit(*deposit2)
	wal.Deposit(*deposit3)

	want := 412.35 + 1350.75 + 20.00 // 1783.10
	got := wal.GetBalanceInRupees()

	assert.Equal(t, got, want)
}

func TestShouldVerifyWithdrawMethodForSameCurrencyWithinBalanceAount(t *testing.T) {
	wal, _ := NewWallet(100.00)
	currency := "INR"

	withdraw1, _ := NewMoney(5.00, currency)
	withdraw2, _ := NewMoney(15.00, currency)
	withdraw3, _ := NewMoney(20.00, currency)

	wal.Withdraw(*withdraw1)
	wal.Withdraw(*withdraw2)
	wal.Withdraw(*withdraw3)

	want := 60.00
	got := wal.GetBalanceInRupees()

	assert.Equal(t, got, want)
}

func TestShouldVerifyWithdrawMethodForDifferentCurrenciesWithinBalanceAount(t *testing.T) {
	wal, _ := NewWallet(10000.00)

	withdraw1, _ := NewMoney(5.00, "USD")
	withdraw2, _ := NewMoney(15.00, "EUR")
	withdraw3, _ := NewMoney(20.00, "USD")

	wal.Withdraw(*withdraw1)
	wal.Withdraw(*withdraw2)
	wal.Withdraw(*withdraw3)

	want := 10000.00 - 412.35 - 1350.75 - 1649.40
	got := wal.GetBalanceInRupees()

	assert.Equal(t, got, want)
}

func TestShouldREturnErrorWithInsufficientBalanceWithDifferentCurrencies(t *testing.T) {
	wal, _ := NewWallet(1000.00)

	withdraw1, _ := NewMoney(5.00, "USD")
	withdraw2, _ := NewMoney(15.00, "EUR")

	wal.Withdraw(*withdraw1)
	err1 := wal.Withdraw(*withdraw2)

	assert.NotNil(t, err1)
}