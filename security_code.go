package main

import "fmt"

type securityCode struct{ code int }

// creating new security code
func newSecurityCode(c int) *securityCode { return &securityCode{c} }

// checking given code matches with account code
func (sc *securityCode) checkCode(code int) error {
	if sc.code != code {
		return fmt.Errorf("given %v code not matched", code)
	}
	fmt.Println("code matched")
	return nil
}
