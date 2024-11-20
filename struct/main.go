package main

import (
	"fmt"
)

func main() {
	invOne := NewInvoice("test")

	invOne.UpdateTip(25)
	invOne.AddItems("item1", 666)
	invOne.AddItems("item2", 888)

	fmt.Println(invOne.format())
}
