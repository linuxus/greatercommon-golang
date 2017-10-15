package main

import (
	"fmt"
)

func main() {
	x := []int{4, 10, 8, 3, 43, 12}
	fmt.Println("This is the original slice \"x\":", x)
	x = append(x, 77, 88, 99, 1024)
	fmt.Println("This is the appended slice \"x\":", x) //Appending more values into the slice

	y := []int{33, 44, 55, 66}                                           //New slice "y"
	x = append(x, y...)                                                  //Appending the whole "y" slice important to note:
	fmt.Println("This is the appended slice \"x\" with slice \"y\":", x) // the "..." is telling func append to add all the value in slice "y"
}
