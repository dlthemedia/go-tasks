// Задание 2: Счётчик с начальным значением
//
// Напиши функцию makeCounter(start int) func() int, которая:
//   - принимает начальное значение
//   - возвращает функцию-счётчик
//   - каждый вызов возвращает следующее число, начиная с start
//
// Ожидаемый вывод:
//   Счётчик от 5:
//   5
//   6
//   7
//   Счётчик от 100:
//   100
//   101
//   Счётчики независимы - счётчик от 5 продолжает:
//   8
//
// Запусти: go run main.go

package main

import (
	"fmt"
)

// TODO: напиши функцию makeCounter(start int) func() int
// Подсказка: текущее значение храни в переменной внутри makeCounter,
// и обращайся к ней из возвращаемой функции (это и есть замыкание)

func makeCounter(start int) func() int {

	current := start

	return func() int {

		current++

		return current

	}

}

func main() {

	fmt.Println("5")
	count1 := makeCounter(5)
	fmt.Println(count1())
	fmt.Println(count1())
	fmt.Println(count1())

	fmt.Println("52")
	count52 := makeCounter(52)
	fmt.Println(count52())
	fmt.Println(count52())
	fmt.Println(count52())

	fmt.Println("Допиши код!")
}
