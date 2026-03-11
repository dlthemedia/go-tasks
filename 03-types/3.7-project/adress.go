package main

type Address struct {
	City string

	Street string
}

type Employee struct {
	Name string

	Address Address
}
