package main

import (
	"fmt"
)

func main() {
	a := []int{43, 23, 3, 6, 89}

	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", a)
}
