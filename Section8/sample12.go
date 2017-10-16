package main

import (
	"fmt"
)

func main() {
	elements := map[string]map[string]string{
		"H": {
			"Name":   "Helium",
			"Symbol": "H",
			"State":  "Gas",
		},
		"Ne": {
			"Name":   "Neon",
			"Symbol": "Ne",
			"State":  "Gas",
		},
		"B": {
			"Name":   "Beryllium",
			"Symbol": "B",
			"State":  "Solid",
		},
	}
	if el, ok := elements["B"]; ok {
		fmt.Println(el["Name"], "state is", el["State"])
	}
}
