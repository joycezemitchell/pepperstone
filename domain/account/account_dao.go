package account

// GetAccountByID ...
func (account *Account) GetAccountByID(id int) (Account, error) {
	rows := GrabAllData()
	for _, row := range rows {
		if row.Id == id {
			return row, nil
		}
	}
	return Account{}, nil
}

// UpdateAccount ...
func (account *Account) UpdateAccount() (error) {
	rows := GrabAllData()
	for _, row := range rows {
		if row.Id == account.Id{
			row.Balance = account.Balance
			return nil
		}
	}
	return nil
}
