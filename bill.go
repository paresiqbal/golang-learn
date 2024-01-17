package main

type bill struct {
	name  string
	items map[string]float64
	tips  float64
}

// make bills
func newBill(name string) bill {
	invoice := bill{
		name:  name,
		items: map[string]float64{},
		tips:  0,
	}

	return invoice
}
