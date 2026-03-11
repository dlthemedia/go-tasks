package main

import (
	"errors"
	"fmt"
	"strconv"
)

type InputError struct {
	Field  string
	Reason string
}

func safeDivide(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("Ты долбоёб на ноль делить?")
	}
	return a / b, nil

}

func createUser(name string) (string, error) {

	if name == "" {

		return "VVP", errors.New("name is empty")

	}
	return name, nil

}

func buyItem(count int) error {

	if count == 0 {
		return ErrOutOfStock
	}
	return nil

}

func readFileMock() error {

	return errors.New("el problema readFileMock")

}

func loadData() error {

	err := readFileMock()

	if err != nil {
		return fmt.Errorf("load data: %w", err)
	}
	return nil
}

func (e InputError) Error() string {
	return e.Field + ": " + e.Reason
}

func validateEmail(email string) error {
	if email == "" {
		return InputError{Field: "email", Reason: "Бля братан ты даже почту указать не можешь?"}
	}
	return nil
}

func parseID(s string) (int, error) {
	if s == "" {
		return 0, errors.New("empty string")
	}

	// если нужно преобразовать строку в число
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return num, nil
}

var ErrOutOfStock = errors.New("out of stock")

func main() {

	fmt.Println(safeDivide(5, 5))

	fmt.Println(safeDivide(5, 0))

	fmt.Println(createUser(""))

	fmt.Println(createUser("Anjela"))

	fmt.Println(buyItem(52))

	fmt.Println(buyItem(0))

	err := readFileMock()
	if err != nil {
		fmt.Println(err)
	}

	err = validateEmail("")
	if err != nil {
		var vErr InputError
		if errors.As(err, &vErr) {
			fmt.Println("validation failed on field:", vErr.Field)

		}
		fmt.Println("other error:", err)
	}

	ids := []string{"52", "69", "", "abc", "6969"}

	for _, id := range ids {
		num, err := parseID(id)
		if err != nil {
			fmt.Println("skip")
			continue
		}
		fmt.Println("ok:", num)
	}

	result, err := Register("Anjela", 52)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Успех:", result)
	}

	result, err = Register("", 52)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Успех:", result)
	}

	result, err = Register("Anjela", -5)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Успех:", result)
	}

}
