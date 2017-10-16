package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":           42,
		"Miss Moneypenny": 27,
		"Abdi":            39,
	}

	fmt.Println(m)
	m["Mike"] = 30
	fmt.Println(m)
	fmt.Println(m["Abdi"])

	delete(m, "James")
	fmt.Println(m)
}
