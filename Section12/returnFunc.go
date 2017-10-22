package main

func main() {

}

func bar() func() int {
	return func() int {
		return 451
	}
}
