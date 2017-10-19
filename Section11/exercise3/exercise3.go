package main

func main() {
	type vehicule struct {
		doors int
		color string
	}
	type truck struct {
		vehicule
		fourWheel bool
	}
	type sedan struct {
		vehicule
		luxury bool
	}
	f150 := truck{

		vehicule: vehicule{
			doors: 2,
			color: "white"
		},
		fourWheel: true,
	}

	focus := sedan{
		vehicule: vehicule{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}
}
