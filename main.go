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

	for key, value := range menu {
		fmt.Println(key, value)
	}

	phoneBook := map[int]string{
		100: "John",
		200: "Paul",
		300: "George",
		400: "Ringo",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[100])

	phoneBook[200] = "Pares"
	fmt.Println(phoneBook)
}
