package main

import (
	transactions "bank/lib/"
	"fmt"
	"time"
)

func main() {

	go func() {
		transactions.Deposit(200)
		fmt.Println("=", transactions.Balance())
	}()

	go transactions.Deposit(500)
	time.Sleep(2 * time.Second)

	fmt.Println("=", transactions.Balance())
}
