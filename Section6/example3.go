package main

import (
	"fmt"
)

func main() {
	x := 0
	for {
		if x > 9 {
			fmt.Println("Existing the main \"For\" Loop")
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("Done!")
}
