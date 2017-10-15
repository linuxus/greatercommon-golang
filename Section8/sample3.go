package main

import (
	"fmt"
)

func main() {
	x := []int{4, 6, 10, 42, 12}
	fmt.Println(x)           //Printing the whole slice
	fmt.Println(x[1:])       //Printing the slice starting at index position 1
	fmt.Println(x[2:len(x)]) //Printing the slice starting at index position 2 and up to the end of the slice
}
