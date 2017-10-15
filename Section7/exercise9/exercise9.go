package main

import (
	"fmt"
)

func main() {
	favSport := "ping-pong"
	switch {
	case favSport == "Baseball":
		fmt.Printf("Fav Sport is %v\n", favSport)
	case favSport == "Basketball":
		fmt.Printf("Fav Sport is %v\n", favSport)
	case favSport == "Hockey":
		fmt.Printf("Fav Sport is %v\n", favSport)
	case favSport == "Soccer":
		fmt.Printf("Fav Sport is %v\n", favSport)
	case favSport == "Football":
		fmt.Printf("Fav Sport is %v\n", favSport)
	default:
		fmt.Println("Your favorite sport NOT in the list")

	}

}
