package main

import (
	"pepperstone/domain/account"
	"pepperstone/services"
	"testing"
)

// TestGetAccount ...
func TestGetAccount(t *testing.T) {

	testAccount := account.Account{
		Id:       12356,
		Balance:  500,
		Currency: "AUD",
	}

	_, err := services.AccountService.GetAccount(testAccount.Id)

	if err != nil {
		t.Errorf("There is an error in GetAccount %v", err)
	}
}

// TestGetAccount ...
func TestAddAmount(t *testing.T) {

	testAccount := account.Account{
		Id:       12356,
		Balance:  500,
		Currency: "AUD",
	}

	testTransfer := account.Transfer{
		From:     343323,
		To:       3243234,
		Amount:   300,
		Currency: "USD",
	}

	err := services.AccountService.AddAmount(testTransfer, testAccount)

	if err != nil {
		t.Errorf("There is an error in AddAmount %v", err)
	}
}


// TestGetAccount ...
func TestDeductAmount(t *testing.T) {

	testAccount := account.Account{
		Id:       12356,
		Balance:  500,
		Currency: "AUD",
	}

	testTransfer := account.Transfer{
		From:     343323,
		To:       3243234,
		Amount:   300,
		Currency: "USD",
	}

	err := services.AccountService.DeductAmount(testTransfer, testAccount)

	if err != nil {
		t.Errorf("There is an error in DeductAmount %v", err)
	}
}
