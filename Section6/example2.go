package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outter loop is at Step: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\tThe inner Loop is at Step: %d\n", j)
		}
	}
}
