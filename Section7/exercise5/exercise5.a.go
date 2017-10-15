package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Printf("Number %v\t has no remainder.\n", i)
		} else if i%4 != 0 {
			fmt.Printf("Number %v\t has the following remainder: %v\n", i, i%4)
		}
	}
}
