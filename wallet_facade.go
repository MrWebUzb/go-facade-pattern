package main

import "log"

type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

func newWalletFacade(accountID string, code int) *walletFacade {
	log.Println("creating new account")
	defer log.Println("Account successfully created")

	return &walletFacade{
		account:      newAccount(accountID),
		wallet:       newWallet(),
		securityCode: newSecurityCode(code),
		notification: &notification{},
		ledger:       &ledger{},
	}
}

func (wf *walletFacade) addMoneyToWallet(accountID string, code int, amount float64) error {
	log.Println("Starting add money to wallet")
	if err := wf.checkAccountAndCode(accountID, code); err != nil {
		return err
	}
	wf.wallet.creditBalance(amount)
	wf.notification.sendWalletCreditNotification(accountID)
	wf.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (wf *walletFacade) deductMoneyFromWallet(accountID string, code int, amount float64) error {
	log.Println("Starting debit money from wallet")
	if err := wf.checkAccountAndCode(accountID, code); err != nil {
		return err
	}
	if err := wf.wallet.debitBalance(amount); err != nil {
		return err
	}
	wf.notification.sendWalletDebitNotification(accountID)
	wf.ledger.makeEntry(accountID, "debit", amount)
	return nil
}

func (wf *walletFacade) checkAccountAndCode(accountID string, code int) error {
	if err := wf.account.checkAccount(accountID); err != nil {
		return err
	}
	if err := wf.securityCode.checkCode(code); err != nil {
		return err
	}
	return nil
}
