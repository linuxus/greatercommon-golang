package main

import (
	"fmt"
)

func main() {

	f := func(x int) {
		fmt.Println("The year big brother start watching:", x)
	}
	f(1984)
}
