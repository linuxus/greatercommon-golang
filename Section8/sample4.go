package main

import (
	"fmt"
)

func main() {
	x := []int{4, 10, 8, 3, 43, 12}
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
