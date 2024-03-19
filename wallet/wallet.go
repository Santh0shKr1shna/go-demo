package wallet

import "errors"

const (
	conversion_rate = 82.47
)

type Wallet struct {
	balance float64
}

func NewWallet(amt float64) (*Wallet, error) {
	if amt < 0.00 {
		return nil, errors.New("amount cannot be negative :(")
	}
	return &Wallet{amt}, nil
}

func (w Wallet) ConvertToDollars() float64 {
	return w.balance / conversion_rate
}

func (w *Wallet) GetBalance() float64 {
	return w.balance
}

func (w *Wallet) Deposit(amt float64) error {
	if amt < 0.00 {
		return errors.New(ErrCannotDepositNegativeAmount)
	}

	w.balance += amt
	return nil
}

func (w *Wallet) Withdraw(amt float64) error {
	if amt < 0.00 {
		return errors.New(ErrCannotWithdrawNegativeAmount)
	}
	if w.balance < amt {
		return errors.New(ErrInsufficientBalance)
	}

	w.balance -= amt
	return nil
}