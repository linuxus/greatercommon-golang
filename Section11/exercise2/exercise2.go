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
	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName, v.lastName)
		fmt.Println("------------------")
		for _, val := range v.favIcecreams {
			fmt.Println(val)
		}
		fmt.Println("====================")
	}

	// fmt.Println("Favorite icecream flavor for ", p1.firstName, p1.lastName, "is:")
	// for _, v := range p1.favIcecreams {
	// 	fmt.Println("\t", v)
	// }
	// fmt.Println("Favorite icecream flavor for ", p2.firstName, p2.lastName, "is:")
	// for _, v := range p2.favIcecreams {
	// 	fmt.Println("\t", v)
	// }
}
