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

	wal, _ := NewWallet(amt)

	assert.Equal(t, 1.00, wal.ConvertToDollars())
}