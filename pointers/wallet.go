package pointers

import "fmt"

// Bitcoin represents Bitcoin currency.
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet represents a wallet containing a Bitcoin currency balance.
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
