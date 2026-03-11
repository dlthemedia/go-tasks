package main

import (
	"errors"
	"strconv"
)

type BankAccount struct {
	Owner string

	Balance int
}

func (account *BankAccount) Deposit(amount int) {

	account.Balance += amount

}

func (account *BankAccount) Withdraw(amount int) error {

	if account.Balance-amount < 0 {
		return errors.New("Ты долбоеб деньги иди зарабатывай")
	}
	account.Balance -= amount
	return nil

}

func (account *BankAccount) Info() string {

	return account.Owner + strconv.Itoa(account.Balance)

}
