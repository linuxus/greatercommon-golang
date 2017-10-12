package main

import (
	"fmt"
)

type (
	tomato string
)

var x tomato
var y string

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = "James Bond"
	fmt.Println(x)
	y = string(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
