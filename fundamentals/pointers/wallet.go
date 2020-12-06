package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin is the currency type
type Bitcoin int

// Wallet object to represent users account
type Wallet struct {
	balance Bitcoin
}

// String formats the output for %s on fmt module
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit add amount to wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance retrives the wallet balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds creates a new error for insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw remove amount from wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
