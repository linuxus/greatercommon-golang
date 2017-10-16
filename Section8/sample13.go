package main

import (
	"fmt"
)

// Program that prints the smallest number in a given array/list
var min int

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	for i, v := range x {
		if i == 0 || v < min {
			min = v
			// fmt.Println("index i:", i)
			// fmt.Println("value v:", v)
			// fmt.Println("value min:", min)
		}

	}
	fmt.Println(min)
}
