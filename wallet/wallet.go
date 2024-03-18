package wallet

import "errors"

const (
	conversion_rate = 82.47
)

type Wallet struct {
	amount float64
}

func NewWallet(amt float64) (*Wallet, error) {
	if amt < 0.00 {
		return nil, errors.New("amount cannot be negative :(")
	}
	return &Wallet{amt}, nil
}

func (w Wallet) ConvertToDollars() float64 {
	return w.amount / conversion_rate
}