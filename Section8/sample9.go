package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":           42,
		"Miss Moneypenny": 32,
		"Batman":          45,
	}

	fmt.Println(m)

	fmt.Println(m["James"])
	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["Batman"])
	fmt.Println(m["Abdillahi"])
	fmt.Println(m["Mike"])

	v, ok := m["Mike"]
	fmt.Println(v, ok)

	if v, ok := m["Mike"]; ok {
		fmt.Println(v, "Does exist!")
	} else {
		fmt.Println(v, "Does NOT Exist!")
	}

}
