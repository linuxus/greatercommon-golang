package main

import (
	"fmt"
)

func main() {
	x := int(42)
	fmt.Printf("x value in decimal: %d\t x value in binary: %b\t x value in Hex: %#x\n", x, x, x)

	y := x << 1
	fmt.Printf("x value in decimal: %d\t x value in binary: %b\t x value in Hex: %#x\n", y, y, y)
}
