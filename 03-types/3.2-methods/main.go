package main

import "fmt"

type Profile struct {
	Name string
}

func printName(p *Profile) {
	if p == nil {
		fmt.Println("profile is nil")
		return
	}

	fmt.Println("name:", p.Name)
}

func main() {
	var p *Profile
	printName(p)

	p = &Profile{Name: "Lena"}
	printName(p)
}
