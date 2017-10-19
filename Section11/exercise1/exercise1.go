package main

import (
	"fmt"
)

func main() {
	type person struct {
		firstName    string
		lastName     string
		favIcecreams []string
	}

	p1 := person{
		firstName: "Abdi",
		lastName:  "Ibrahim",
		favIcecreams: []string{
			"Vanilla",
			"Chocolate",
			"GreenTea",
		},
	}

	p2 := person{
		firstName: "Mohamad",
		lastName:  "Ali",
		favIcecreams: []string{
			"Peach",
			"Tiramisu",
			"Coffee",
		},
	}

	fmt.Println("Favorite icecream flavor for ", p1.firstName, p1.lastName, "is:")
	for _, v := range p1.favIcecreams {
		fmt.Println("\t", v)
	}
	fmt.Println("Favorite icecream flavor for ", p2.firstName, p2.lastName, "is:")
	for _, v := range p2.favIcecreams {
		fmt.Println("\t", v)
	}
}
