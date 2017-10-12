package main

import (
	"fmt"
)

type (
	potato int
	tomato string
	onion  bool
)

var x potato
var y tomato
var z onion

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	y = "James Bond"
	fmt.Println(y)
	z = false
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	z = true
	fmt.Println(z)
}
