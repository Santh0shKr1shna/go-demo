package wallet

import (
	"errors"
	"math"
)

type Rupee float64

type conversion_rate float64

const (
	USD conversion_rate = 82.47
	EUR conversion_rate = 90.05
)

var supported = []string{"INR", "USD", "EUR"}
var currrencyMap = map[string]conversion_rate{
	"USD": USD,
	"EUR": EUR,
}

// balance is always in INR
type Wallet struct {
	balance Rupee
}

type Money struct {
	value float64
	currrency string
}

func NewMoney(amt float64, cur string) (*Money, error) {
	if amt <= 0.00 {
		return nil, errors.New(ErrInitialAmountNegative)
	}
	if !isAvailable(supported, cur) {
		return nil, errors.New(ErrCurrencyNotSupported)
	}

	return &Money{
		value: amt,
		currrency: cur,
	}, nil
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

func (w Wallet) ConvertToRupee(m Money) float64 {
	if m.currrency == "INR" {
		return m.value
	}

	rate := currrencyMap[m.currrency]
	return m.value * float64(rate)
}

func (w Wallet) GetBalanceInRupees() float64 {
	return toFixed(float64(w.balance), 2)
}

func (w Wallet) GetBalance(cur string) float64 {
	rate := currrencyMap[cur]
	amt := toFixed(
		float64(w.balance) / float64(rate),
		2,
	)

	return amt
}

func (w *Wallet) Deposit(m Money)  {
	if m.currrency == "INR" {
		w.balance += Rupee(m.value)
		return 
	}
	w.balance += Rupee(w.ConvertToRupee(m))
}

func (w *Wallet) Withdraw(m Money) error {
	rate := currrencyMap[m.currrency]
	amt := m.value * float64(rate)

	if amt > w.GetBalanceInRupees() {
		return errors.New(ErrInsufficientBalance)
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

func isAvailable(slice []string, cur string) bool {
	for _, c := range slice {
		if c == cur {
			return true
		}
	}
	return false
}