package main

import (
	"errors"
	"fmt"
)

type PaymentProcessor interface {
	Pay(amount int) error
}

type CardProcessor struct {
	CardNumber string
}

func (c CardProcessor) Pay(amount int) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	fmt.Printf("Оплачено %d с карты %s\n", amount, c.CardNumber)
	return nil
}

type CashProcessor struct{}

func (c CashProcessor) Pay(amount int) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	fmt.Printf("Оплачено %d наличными\n", amount)
	return nil
}

func checkout(p PaymentProcessor, amount int) {
	err := p.Pay(amount)
	if err != nil {
		fmt.Println("Ошибка оплаты:", err)
	}
}
