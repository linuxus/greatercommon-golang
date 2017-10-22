package main

import "fmt"

func main() {
	foo()
	func() {
		fmt.Println("Anonymous func ran")
	}()
	func(x int) {
		fmt.Println("Anonymous func with parameter", x, "ran")
	}(42)
}

func foo() {
	fmt.Println("foo func ran")
}
