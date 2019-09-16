package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErrInsufficientFunds = errors.New("Insufficient funds to withdraw")

type Wallet struct {
	balance Bitcoin
}

//this interface is defined in the fmt package and lets you define how your type is printed when used with the %s format string in prints
type Stringer interface {
	String() string
}

//in Go, when you call a function or a method the arguments are copied
//if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds

	}
	w.balance -= amount
	return nil

}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
