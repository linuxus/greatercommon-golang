package main

import (
	"fmt"
)

func main() {
	n2 := fact(4)
	fmt.Println(n2)

}

func fact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n

	}
	return total
}
