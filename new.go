package main

import (
	"fmt"
	"math"
)

func cycleName(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func greet(n string) {
	fmt.Printf("Hello %v \n", n)
}

func circle(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	cycleName([]string{"Pares", "Paul", "Paulo"}, greet)

	a1 := circle(10.5)
	a2 := circle(15)

	fmt.Printf()
	fmt.Println(a1, a2)
}
