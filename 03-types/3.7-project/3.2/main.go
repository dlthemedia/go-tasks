package main

import "fmt"

func resetScore(score *int) {
	*score = 0
}

func printUserName(u *User) {
	if u == nil {
		fmt.Println("user is nil")
	} else {
		fmt.Println(u.Name)
	}
}

func main() {

	wallet := Wallet{}

	wallet.Deposit(100)

	wallet.Deposit(50)

	fmt.Printf("Итоговый баланс: %d\n", wallet.Balance)

	wallet = Wallet{}

	wallet.DepositCopy(100)

	wallet.DepositCopy(50)

	fmt.Printf("Итоговый баланс: %d\n", wallet.Balance)

	score := 100
	fmt.Println("До вызова resetScore:", score)

	resetScore(&score)

	fmt.Println("После вызова resetScore:", score)

	var u1 *User
	printUserName(u1)

	u2 := &User{Name: "Alice"}
	printUserName(u2)

	acc := BankAccount{
		Owner:   "Alice",
		Balance: 1000,
	}

	// Выполняем операции
	acc.Deposit(500)
	fmt.Println(acc.Withdraw(200))
	fmt.Println(acc.Withdraw(1500)) // здесь будет ошибка

	// Выводим финальное состояние
	fmt.Println(acc.Info())

}
