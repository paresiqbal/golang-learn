package main

import "fmt"

func greet(n string) {
	fmt.Printf("Hello %v \n", n)
}

func cycleName(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func main() {
	greet("Pares")
	cycleName([]string{"Pares", "Paul", "Paulo"}, greet)
}
