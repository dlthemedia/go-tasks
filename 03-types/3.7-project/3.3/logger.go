package main

import (
	"fmt"
	"time"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

type PrefixLogger struct {
	Prefix string
}

func (l PrefixLogger) Log(message string) {
	fmt.Println(l.Prefix + message)
}

func processOrder(logger Logger, id string) {
	logger.Log("Начало обработки заказа " + id)
	logger.Log("Проверка данных заказа...")
	time.Sleep(500000000)
	logger.Log("Заказ " + id + " успешно обработан")
}
