package main

import (
	"fmt"
)

func main() {
	var x [5]int
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}
