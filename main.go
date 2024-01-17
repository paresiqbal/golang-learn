package main

import "fmt"

func updateName(x string) string {
	x = "Pares"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffe"] = 3.99
}

func main() {
	name := "Risti"
	name = updateName(name)

	fmt.Println(name)

	menu := map[string]float64{
		"Pizza": 10.0,
		"Pasta": 15.0,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
