package csvprocess

import (
	"pepperstone/domain/account"
	"pepperstone/services"
)

// Process Cash Transfer Type CSV ...
func ProcesCTT(file string, ch chan bool) error {

	var CSVoutData []account.Transfer

	// Read csv file
	data, err := ReadCSV(file)
	if err != nil {
		return err
	}

	// Convert csv data to account.Transfer
	rows, _ := GetData(data)

	for _, row := range rows {
		errorTransfer := false
		errorMessage := ""
		var status string

		accountFrom, _ := services.AccountService.GetAccount(row.From) // Get account information for sender
		accountTo, _ := services.AccountService.GetAccount(row.To)     // Get account information for receiver

		// Validate if the from currency is equal to destination currency, otherwise, do not process that record
		if row.Currency != accountTo.Currency {
			errorMessage = "Destination account currency mismatch"
			errorTransfer = true
			status = "failed"
		}

		// Validate to make sure the amount is not bigger than the from account balance (no negative balance).
		if accountFrom.Balance <= row.Amount {
			errorMessage = "Sender does not have enough balance to send money"
			errorTransfer = true
			status = "failed"
		}

		if !errorTransfer {
			services.AccountService.AddAmount(row, accountTo)    // Update balance of the reciever
			services.AccountService.DeductAmount(row, accountFrom) // Update balance of the sender
		}

		CSVoutData = append(CSVoutData, account.Transfer{
			From:     row.From,
			To:       row.To,
			Amount:   row.Amount,
			Currency: row.Currency,
			Status:   status,
			Error:    errorMessage,
		})

		// Generate CSV
		GenerateCSV(file, CSVoutData)
	}
	ch <- true
	return nil
}
