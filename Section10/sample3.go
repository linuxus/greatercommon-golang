package main

import (
	"fmt"
)

func main() {
	sa := struct {
		first string
		last  string
		age   int
		ltk   bool
	}{
		first: "James",
		last:  "Bond",
		age:   42,
		ltk:   true,
	}

	fmt.Println(sa.first, sa.last, sa.age, sa.ltk)
}
