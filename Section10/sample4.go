package main

import (
	"fmt"
)

func main() {
	type person struct {
		first string
		last  string
		age   int
	}

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   42,
	}

	type secretAgent struct {
		person
		ltk bool
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   45,
		},
		ltk: true,
	}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
}
