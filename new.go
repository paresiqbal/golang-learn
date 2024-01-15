package main

import (
	"fmt"
	"strings"
)

func getInit(name string) (string, string) {
	s := strings.ToUpper(name)
	names := strings.Split(s, " ")

	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	firstName, lastName := getInit("Pares iqbal")
	fmt.Println(firstName, lastName)
}
