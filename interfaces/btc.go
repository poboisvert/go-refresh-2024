package main

import "errors"

type BitcoinAccount struct {
	name    string
	balance int
	fee     int
}

func NewBitcoinAccount() *BitcoinAccount {
	return &BitcoinAccount{
		name:    "bitcoin",
		balance: 0,
		fee:     100,
	}
}

func (b *BitcoinAccount) GetAccountName() string {
	return b.name
}

func (b *BitcoinAccount) GetBalance() int {
	return b.balance
}

func (b *BitcoinAccount) Deposit(amount int) {
	b.balance += amount
}

func (b *BitcoinAccount) Withdraw(amount int) error {
	newBalance := b.balance - amount - b.fee
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}
	b.balance = newBalance
	return nil
}
