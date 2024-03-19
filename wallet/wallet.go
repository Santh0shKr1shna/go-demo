package wallet

import (
	"errors"
	"math"
)

type Rupee float64

type conversion_rate float64

const (
	USD conversion_rate = 82.47
)

// balance is always in INR
type Wallet struct {
	balance Rupee
}

func NewWallet(amt float64) (*Wallet, error) {
	if amt < 0.00 {
		return nil, errors.New(ErrInitialAmountNegative)
	}
	return &Wallet{Rupee(amt)}, nil
}

func (w Wallet) ConvertToDollars(amt float64) float64 {
	roundup := toFixed(amt / float64(USD), 2)
	return roundup
}

func (w Wallet) GetBalance() float64 {
	return toFixed(float64(w.balance), 2)
}

func (w Wallet) GetBalanceInDollars() float64 {
	return w.ConvertToDollars(float64(w.balance))
}

func (w *Wallet) Deposit(amt float64, cur string) error {
	if amt < 0.00 {
		return errors.New(ErrCannotDepositNegativeAmount)
	}
	if cur != "INR" && cur != "USD" {
		return errors.New(ErrCurrencyNotSupported)
	}

	if cur == "INR" {
		w.balance += Rupee(amt)
		return nil
	}

	w.balance += Rupee(w.ConvertToDollars(amt))
	return nil
}

func (w *Wallet) Withdraw(amt float64, cur string) error {
	if amt < 0.00 {
		return errors.New(ErrCannotWithdrawNegativeAmount)
	}
	if float64(w.balance) < amt {
		return errors.New(ErrInsufficientBalance)
	}
	if cur != "INR" && cur != "USD" {
		return errors.New(ErrCurrencyNotSupported)
	}

	if cur == "USD" {
		if float64(w.balance) < w.ConvertToDollars(amt) {
			return errors.New(ErrInsufficientBalance)
		}
		w.balance -= Rupee(w.ConvertToDollars(amt))
		return nil
	}
	w.balance -= Rupee(amt)
	return nil
}

// UTILITIES

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}