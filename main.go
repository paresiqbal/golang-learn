package main

import "fmt"

func main() {
	myBill := newBill("Pares invoice:")

	myBill.formatBill()
	myBill.addItem("Coffee", 10000)
	myBill.updateTip(10)

	fmt.Println(myBill.formatBill())
}
