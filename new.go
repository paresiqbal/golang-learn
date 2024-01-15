package main

import "fmt"

func main() {
	name := []string{"Peter", "Paul", "Mary"}

	for index, value := range name {
		fmt.Printf("position at index %v is %v", index, value)
	}
}
