package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Chocolate", "Vanila", "Martini"}
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut", "Peach"}
	fmt.Println(jb)
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
