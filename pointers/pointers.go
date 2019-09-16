package pointers

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

//this interface is defined in the fmt package and lets you define how your type is printed when used with the %s format string in prints
type Stringer interface {
	String() string
}

//in Go, when you call a function or a method the arguments are copied
//because of that, we need to use a pointer of real wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
