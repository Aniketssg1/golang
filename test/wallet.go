package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(deposit Bitcoin) {
	w.balance += deposit
	fmt.Printf("Deposited amount: %v\n", deposit)
	fmt.Printf("Current Balance is: %v\n", w.balance)
}

func (w *Wallet) Withdraw(withdraw Bitcoin) error {
	if w.balance < withdraw {
		return errors.New("auukaat ke bahar mt ja...:)")
	} else {
		w.balance -= withdraw
	}
	fmt.Printf("Withdrawn amount: %v\n", withdraw)
	fmt.Printf("Current Balance is: %v\n", w.balance)
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("Balance: %v\n", w.balance)
	return (*w).balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
