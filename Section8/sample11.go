package main

import (
	"fmt"
)

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Nu"])
}
