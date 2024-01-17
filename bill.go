package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make bills
func newBill(name string) bill {
	invoice := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
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

	// add tip
	fs += fmt.Sprintf("%v Rp...%0.2f", "tip:", invoice.tip)

	// total bill
	fs += fmt.Sprintf("%v Rp...%0.2f", "total:", total+invoice.tip)

	return fs
}

// update tip
func (invoice *bill) updateTip(tip float64) {
	invoice.tip = tip
}

// add item to bill
func (invoice *bill) addItem(name string, price float64) {
	invoice.items[name] = price
}
