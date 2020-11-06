package main

import "log"

func main() {
	wf := newWalletFacade("asliddin", 12345)

	if err := wf.addMoneyToWallet("asliddin", 12345, 500.0); err != nil {
		log.Printf("cannot add money: %v\n", err)
	}
	if err := wf.deductMoneyFromWallet("asliddin", 12345, 300.0); err != nil {
		log.Printf("cannot debit from wallet: %v\n", err)
	}
}
