package main

import (
	"fmt"
)

func main() {
	a := incrementor()
	b := incrementor()
	c := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

	for i := 0; i <= 9; i++ {
		fmt.Println("index:", i, c())
	}
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
