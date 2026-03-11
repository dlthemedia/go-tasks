package main

import "fmt"

type Greeter interface {
	Greet() string
}

type User struct {
	Name string
}

func (u User) Greet() string {
	return "Чикса я уже смешарик " + u.Name
}

type Guest struct {
	SessionID string
}

func (g Guest) Greet() string {
	return "Я нахуй гость " + g.SessionID
}

func PrintGreeting(g Greeter) {
	fmt.Println(g.Greet())
}
