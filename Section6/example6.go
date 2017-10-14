package main

import (
	"fmt"
)

func main() {
	for i := 32; i < 123; i++ {
		fmt.Printf("Decimal value of \"i\": %d\t its Unicode value %U\t The character represented by its Unicode %c\n", i, i, i)
	}

}
