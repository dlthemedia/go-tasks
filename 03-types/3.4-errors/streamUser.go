package main

import (
	"errors"
	"fmt"
)

func validateName(name string) error {

	if name == "" {
		return errors.New("Братан ты долбоеб пустое имя ставить?")
	}
	return nil
}

func validateAge(age int) error {

	if age < 0 {
		return errors.New("Ай сын ша машалла красавец меньше нуля возраст ставить")
	}

	return nil

}

func Register(name string, age int) (string, error) {
	if err := validateName(name); err != nil {
		return "", fmt.Errorf("register: неверное имя: %w", err)
	}

	if err := validateAge(age); err != nil {
		return "", fmt.Errorf("register: неверный возраст: %w", err)
	}

	return "Регистрация да", nil
}
