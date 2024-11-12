package main

import "errors"

type WellsFargo struct {
	name    string
	balance int
}

func NewWellsFargo() *WellsFargo {
	return &WellsFargo{
		name:    "wellsfargo",
		balance: 0,
	}
}

func (w *WellsFargo) GetAccountName() string {
	return w.name
}

func (w *WellsFargo) GetBalance() int {
	return w.balance
}

func (w *WellsFargo) Deposit(amount int) {
	w.balance += amount
}

func (w *WellsFargo) Withdraw(amount int) error {
	newBalance := w.balance - amount
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}
	w.balance = newBalance
	return nil
}
