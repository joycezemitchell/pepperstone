package services

import (
	"pepperstone/domain/account"
)

var (
	AccountService accountServiceInterface = &accountService{}
)

type accountService struct{}

type accountServiceInterface interface {
	AddAmount(transfer account.Transfer, account account.Account) error
	DeductAmount(transfer account.Transfer, account account.Account) error
	GetAccount(id int)  (account.Account, error)
}

// AddAmount ...
func (a *accountService) AddAmount(transfer account.Transfer, account account.Account) error {
	account.Balance = account.Balance + transfer.Amount
	account.UpdateAccount()
	return nil
}

// DeductAmount ...
func (a *accountService) DeductAmount(transfer account.Transfer,  account account.Account) error {
	account.Balance = account.Balance - transfer.Amount
	account.UpdateAccount()
	return nil
}

// GetAccount ...
func (a *accountService) GetAccount(id int) (account.Account, error) {
	var act account.Account
	data, _ := act.GetAccountByID(id)
	return data, nil
}
