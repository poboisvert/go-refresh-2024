package main

import (
	operations "bank/lib"
	"fmt"
	"time"
)

func main() {

	// No memory if outside go goroutine
	//go func() {
	//	operations.Deposit(200)
	//	fmt.Println("=", operations.Balance())
	//}()

	operations.Deposit(300)
	operations.Deposit(300)
	operations.Withdraw(400)

	time.Sleep(2 * time.Second)

	fmt.Println("Balance:", operations.Balance())
}
