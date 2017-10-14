package main

import (
	"fmt"
)

const (
	x uint64  = 42
	y float64 = 12.040585
	z         = "James Bond"
)

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
