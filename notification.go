package main

import "log"

type notification struct{}

func (n *notification) sendWalletCreditNotification(ID string) {
	log.Printf("sending wallet credit notification for account %s\n", ID)
}
func (n *notification) sendWalletDebitNotification(ID string) {
	log.Printf("sending wallet credit notification for account %s\n", ID)
}
