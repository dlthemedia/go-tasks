package main

type Wallet struct {
	Balance int
}

func (w *Wallet) Deposit(amount int) {
	w.Balance += amount
}

func (w Wallet) DepositCopy(amount int) {
	w.Balance += amount
}
