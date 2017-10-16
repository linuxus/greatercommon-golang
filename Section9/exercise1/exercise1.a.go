package main

import "fmt"

func main() {
	var a [5]int
	a[0] = 10
	a[1] = 3
	a[1] = 72
	a[1] = 98
	a[1] = 43

	for i := range a {
		fmt.Println(a[i])
	}
	fmt.Printf("%T\n", a)
}
