package main

import (
	"fmt"
)

func main() {
	x := make([]int, 5, 10)
	//x = []int{4, 10, 8, 3, 43, 12}
	x[0] = 3
	x[2] = 42
	x[4] = 99
	x = append(x, 3020, 3021, 3022, 3023, 3024, 3025)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
