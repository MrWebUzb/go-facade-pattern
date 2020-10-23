package main

import (
	"fmt"
	"log"
)

type account struct{ ID string }

func newAccount(ID string) *account {
	return &account{ID}
}

func (a *account) checkAccount(ID string) error {
	log.Printf("checking account with %s id\n", ID)
	if a.ID != ID {
		return fmt.Errorf("account ID not verified")
	}
	log.Printf("account verified")
	return nil
}
