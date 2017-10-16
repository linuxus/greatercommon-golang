package main

import (
	"fmt"
)

func main() {
	s := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	for i, r := range s {
		fmt.Println(i, r)
		for rg, d := range s[i] {
			fmt.Println(rg, d)
		}
	}
}
