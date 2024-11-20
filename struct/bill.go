package main

import "fmt"

type invoice struct {
	name  string
	items map[string]float64
	tip   float64
}

func NewInvoice(name string) invoice {
	i := invoice{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return i
}

// w/ pointer && formatting
func (i *invoice) format() string {
	fs := fmt.Sprintf("Invoice for: %s\nTip: $%0.2f\nInvoice breakdown: \n", i.name, i.tip)
	var total float64 = 0

	for k, v := range i.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}

// w/ pointer so we do not do a copy
func (i *invoice) UpdateTip(tip float64) {
	i.tip = tip
}

func (i *invoice) AddItems(key string, price float64) {
	i.items[key] = price
}
