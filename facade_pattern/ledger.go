package main

import "log"

type ledger struct{}

func (l *ledger) makeEntry(ID string, t string, a float64) {
	log.Printf("make ledger entry for account %s with %s type for amount %v", ID, t, a)
}
