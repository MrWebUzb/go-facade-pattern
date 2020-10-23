package main

import (
	"fmt"
	"log"
)

type wallet struct{ balance float64 }

func newWallet() *wallet { return &wallet{} }

func (w *wallet) creditBalance(amount float64) {
	w.balance += amount
	log.Println("wallet balance added successfully")
}

func (w *wallet) debitBalance(amount float64) error {
	if w.balance < amount {
		return fmt.Errorf("is not sufficient balance in wallet")
	}
	w.balance -= amount
	log.Println("successfully debited")
	return nil
}
