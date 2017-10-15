package main

import (
	"fmt"
)

func main() {
	//Creating a slice
	x := []int{4, 2, 42, 65, 10, 11, 0}
	fmt.Println(x)
	fmt.Println(len(x))

	//Using range to loop through the slice
	for i, v := range x {
		fmt.Println(i, v) //Printing the index and value in the slice
	}
	for _, v := range x {
		fmt.Println(v) //Printing just the value and throwing away the index
	}
}
