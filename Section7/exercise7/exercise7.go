package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("Case 1 is false so... prints")
	case 2 == 2:
		fmt.Println("Case 2 true so... prints")
	case 7 == 11:
		fmt.Println("Case 3 is false so... no prints")
	default:
		fmt.Println("No other cases matched!")
	}

}
