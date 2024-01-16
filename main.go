package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":    10.0,
		"noodles": 15.0,
		"rice":    20.0,
		"curry":   25.0,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])
}
