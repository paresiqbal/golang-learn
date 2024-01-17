package main

import "fmt"

func main() {
	myBill := newBill("Pares invoice:")

	myBill.formatBill()

	fmt.Println(myBill.formatBill())
}
