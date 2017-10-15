package main

import (
	"fmt"
)

func main() {
	name := "CI"
	if name == "AI" {
		fmt.Println("Name is AI!")
	} else if name == "BI" {
		fmt.Println("Name is BI!")
	} else {
		fmt.Println("Name Not AI nor BI... Please put the correct Name in variable \"name\"")
	}
}
