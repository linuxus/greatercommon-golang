package main

import (
	"fmt"
)

func main() {

	var p = struct {
		first      string
		last       string
		age        int
		favFlavors []string
		address    map[string]string
	}{
		first: "James",
		last:  "Bond",
		age:   43,
		favFlavors: []string{
			"Vanilla",
			"Chocolate",
			"Coffee",
		},
		address: map[string]string{
			"Street Name": "Bethany St.",
			"City":        "Markham",
			"Country":     "Canada",
			"Postal Code": "M1S 4T9",
		},
	}

	fmt.Println(p.first, p.last, p.age)
	fmt.Println("Favorite IceCream Flavors:")
	fmt.Println("--------------------")
	for _, v := range p.favFlavors {
		fmt.Println("\t", v)
	}
	fmt.Println("And hist address is:")
	fmt.Println("--------------------")
	for _, v := range p.address {
		fmt.Println("\t", v)
	}
}
