package main

import "fmt"

type ContactInfo struct {
	Phone string
	Email string
}

type Address struct {
	City string
	Zip  string
}

type Order struct {
	ID      string
	Address Address
	ContactInfo
}

func main() {
	order := Order{
		ID: "ORD-1001",
		Address: Address{
			City: "Moscow",
			Zip:  "101000",
		},
		ContactInfo: ContactInfo{
			Phone: "+79990000000",
			Email: "buyer@example.com",
		},
	}

	fmt.Println("Order:", order.ID)
	fmt.Println("City:", order.Address.City)
	fmt.Println("Email:", order.Email)
}
