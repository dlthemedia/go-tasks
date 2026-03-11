// Задание 2: Идентификатор запроса через контекст
//
// Нужно передать идентификатор запроса через цепочку функций с помощью context.WithValue.
//
// 1. Создай свой тип для ключа контекста:
//    type contextKey string
//    const requestIDKey contextKey = "request-id"
//
// 2. Напиши функцию middleware(next func(ctx context.Context)):
//    - создаёт контекст с идентификатором "req-42" через context.WithValue
//    - вызывает next(ctx) передавая обогащённый контекст
//
// 3. Напиши функцию handler(ctx context.Context):
//    - достаёт идентификатор из контекста через ctx.Value(requestIDKey)
//    - выводит: "Обрабатываем запрос: req-42"
//    - если идентификатора нет - выводит предупреждение
//
// Ожидаемый вывод:
//   Обрабатываем запрос: req-42
//
// Запусти: go run main.go

package main

import (
	"context"
	"fmt"
)

type contextKey string

const requestIDKey contextKey = "request-id"

func middleware(next func(ctx context.Context)) {

	ctx := context.WithValue(context.Background(), requestIDKey, "req-42")
	next(ctx)
}

func handler(ctx context.Context) {

	reqID := ctx.Value(requestIDKey)
	if reqID == nil {
		fmt.Println("Предупреждение: идентификатор запроса не найден")
		return
	}

	fmt.Printf("Обрабатываем запрос: %s\n", reqID)
}

func main() {
	middleware(handler)
}
