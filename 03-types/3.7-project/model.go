package main

type Package struct {
	ID     string
	Weight int
}

type Destination struct {
	City string
	Zip  string
}

type Shipment struct {
	Package     Package
	Destination Destination
}
