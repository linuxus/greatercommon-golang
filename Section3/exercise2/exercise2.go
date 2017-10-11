package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Println(x, " | ", y, " | ", z)
	//Printing variables default formats
	fmt.Printf("%v\t%v\t%v\n", x, y, z)
	//Printing variables TYPE
	fmt.Printf("%T\t%T\t%T", x, y, z)
}
