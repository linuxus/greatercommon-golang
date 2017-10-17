package main

import (
	"fmt"
)

func main() {

	type person struct {
		firstName string
		lastName  string
		age       int
		country   string
		city      string
		phone     string
		id        int
	}
	p1 := person{
		"Abdi",
		"Ibrahim",
		39,
		"Canada",
		"Toronto",
		"416-509-4565",
		12345,
	}
	p2 := person{

		age:       42,
		country:   "United Kingdom",
		city:      "London",
		phone:     "708- 6509-4565",
		firstName: "James",
		lastName:  "Bond",
	}

	fmt.Println(p1.firstName, p1.phone, p1.id)
	fmt.Println(p2.firstName)
}
