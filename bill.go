package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tips  float64
}

// make bills
func newBill(name string) bill {
	invoice := bill{
		name:  name,
		items: map[string]float64{"item1": 100.2, "item2": 200.7},
		tips:  0,
	}

	return invoice
}

// format the bill
func (invoice bill) formatBill() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for key, value := range invoice.items {
		fs += fmt.Sprintf("%v  Rp...%v \n", key+":", value)
		total += value
	}

	// total bill
	fs += fmt.Sprintf("%v Rp...%0.2f", "total:", total)

	return fs
}
