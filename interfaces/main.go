package main

import (
	"fmt"
)

// IBankAccount interface
type IBankAccount interface {
	GetAccountName() string
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {

	myAccounts := []IBankAccount{
		NewWellsFargo(),
		NewBitcoinAccount(),
	}

	for _, account := range myAccounts {
		balance := account.GetBalance()
		fmt.Printf("[%s] default balance = %d\n", account.GetAccountName(), balance)

		account.Deposit(2000)
		if err := account.Withdraw(600); err != nil {
			fmt.Printf("account.Withdraw(600) fail: %v", err)
		}

		balance = account.GetBalance()
		fmt.Printf("[%s] balance = %d\n", account.GetAccountName(), balance)
	}
}

// RUN: go run . to make it work with all files
