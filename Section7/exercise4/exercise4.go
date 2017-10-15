package main

import "fmt"

func main() {
	bd := 1978
	for {
		if bd > 2017 {
			break
		} else {
			fmt.Println(bd)
			bd++
		}

	}
}
