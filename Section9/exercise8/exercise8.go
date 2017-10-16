package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"bond_james": {
			"Shaken",
			"not stirred",
			"Martinis",
			"Women",
		},
		"moneypenny_miss": {
			"James Bond",
			"Literature",
			"Computer Science",
		},
		"no_dr": {
			"Being evil",
			"Ice cream",
			"Sunsets",
		},
	}

	m["Abdi"] = []string{"Cloud", "SDN", "Golang"}
	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println("This is for", k)
		for i, val := range v {
			//fmt.Println("This is the index for", i)
			fmt.Println("\t Favorit item:", val, "| it's index position:", i)
		}
	}
}
