package main

import (
	"fmt"
)

func main() {
	month := "August"
	switch {
	case month == "August":
		fmt.Println("This is August")

	case month == "September":
		fmt.Println("This is September")

	case month == "October":
		fmt.Println("This is October")

	case month == "November":
		fmt.Println("This is November")
	}

}
