package main

import "fmt"

func printStringLength(x any) {
	if s, ok := x.(string); ok {
		fmt.Println(len(s))
	} else {
		fmt.Println("not a string")
	}
}

func main() {
	user := User{Name: "Alice"}
	guest := Guest{SessionID: "abc123"}

	PrintGreeting(user)
	PrintGreeting(guest)

	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 4},
		Rectangle{Width: 3.5, Height: 2},
		Circle{Radius: 7.2},
	}

	for i, shape := range shapes {
		fmt.Printf("Фигура %d: площадь = %.2f\n", i+1, shape.Area())
	}

	a := Athlete{Name: "Usain"}
	fmt.Println(a.Run())

	var eng Engine
	runCycle(eng)

	printStringLength("hello")
	printStringLength(42)
	printStringLength([]int{1, 2, 3})

	card := CardProcessor{CardNumber: "1234-5678-9012-3456"}
	cash := CashProcessor{}

	checkout(card, 1000)
	checkout(cash, 500)
	checkout(card, -50)

	fmt.Println("=== ConsoleLogger ===")
	console := ConsoleLogger{}
	processOrder(console, "ORD-123")

	fmt.Println("\n=== PrefixLogger ===")
	prefixed := PrefixLogger{Prefix: "[INFO] "}
	processOrder(prefixed, "ORD-456")

}
