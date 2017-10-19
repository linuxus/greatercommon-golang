package main

import (
	"fmt"
)

func main() {

	type person struct {
		first      string
		last       string
		favFlavors []string
	}

	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"Vanilla",
			"Chocolate",
			"Tiramisu",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"Strawberry",
			"Peach",
			"COffee",
		},
	}

	fmt.Println(p1.first, p1.last)
	fmt.Println("==================")
	for k, v := range p1.favFlavors {
		fmt.Println(k, v)
	}

	fmt.Println(p2.first, p2.last)
	fmt.Println("==================")
	for k, v := range p2.favFlavors {
		fmt.Println(k, v)
	}
}
