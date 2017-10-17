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

	type secretAgent struct {
		person
		id    int
		ltk   bool
		first string
	}
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   42,
		},
		id:    12345,
		ltk:   true,
		first: `007`,
	}

	fmt.Println(sa1.first, sa1.person.first, sa1.last, sa1.age, sa1.id, sa1.ltk)
}
