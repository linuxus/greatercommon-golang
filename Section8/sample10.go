package main

import "fmt"

func main() {
	var x = make(map[string]int)
	x["key1"] = 1
	x["key2"] = 2
	x["key3"] = 3
	x["key4"] = 4
	x["key5"] = 5
	x["key6"] = 6
	fmt.Println(x)
	//Delete key5
	delete(x, "key5")
	fmt.Println(len(x))

}
