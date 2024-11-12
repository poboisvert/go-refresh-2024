package main

import "fmt"

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

		account.Deposit(1000)
		if err := account.Withdraw(400); err != nil {
			fmt.Printf("account.Withdraw(400) fail: %v", err)
		}

		balance = account.GetBalance()
		fmt.Printf("[%s] balance = %d\n", account.GetAccountName(), balance)
	}

}
