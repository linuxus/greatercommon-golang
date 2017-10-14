package main

import (
	"fmt"
)

func main() {
	for x := 0; x <= 100; x++ {
		if x == 0 {
			fmt.Printf("Value \"x\": %d\t is Even Number\n", x)
		} else if x%2 == 0 {
			fmt.Printf("Value \"x\": %d\t is Even Number\n", x)
		} else {
			fmt.Printf("Value \"x\": %d\t is Odd Number\n", x)
		}
	}
}
