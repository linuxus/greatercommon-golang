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
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first, v.last)
		fmt.Println("==============")
		for _, val := range v.favFlavors {
			fmt.Println(val)
		}

	}

}
