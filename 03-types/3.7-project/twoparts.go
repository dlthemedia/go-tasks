package main

type ContactInfo struct {
	Phone string
	Email string
}

type AddresAddress struct {
	City   string
	Street string
}

type Client struct {
	ID            string
	AddresAddress Address
	ContactInfo
}
