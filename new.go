package main

import "fmt"

func greet(n string) {
	fmt.Printf("Hello %v \n", n)
}

func bye(n string) {
	fmt.Printf("Bye %v \n", n)
}

func main() {
	greet("Pares")
	bye("Pares")
}
