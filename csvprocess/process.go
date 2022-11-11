package csvprocess

import (
	"encoding/csv"
	"fmt"
	"os"
	"pepperstone/domain/account"
	"strconv"
)


// ReadCSV ...
func ReadCSV(file string) ([][]string, error) {

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(dir + "/csv/input/" + file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}

// GetData ...
func GetData(data [][]string) ([]account.Transfer, error) {

	var dataTransfer []account.Transfer
	for _, line := range data {
		from, _ := strconv.Atoi(line[0])
		to, _ := strconv.Atoi(line[1])
		amount, _ := strconv.Atoi(line[2])
		dataTransfer = append(dataTransfer, account.Transfer{
			From:     from,
			To:       to,
			Amount:   amount,
			Currency: line[3],
		})
	}
	return dataTransfer, nil
}

// GenerateCSV ...
func GenerateCSV(filename string, transfer []account.Transfer) error{
	dir, err := os.Getwd()
	if err != nil {
		return err
	}


	file, err := os.Create(dir + "/csv/output/" + filename)
	if err != nil {
		return err
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	var data [][]string
	data = append(data, []string{
		"From Account(id)",
		"To Account(id)",
		"Amount",
		"Currency",
		"Status",
		"Error",
	})


	for _, record := range transfer {
		row := []string{
			strconv.Itoa(record.From),
			strconv.Itoa(record.To),
			strconv.Itoa(record.Amount),
			record.Currency,
			record.Status,
			record.Error,
		}
		data = append(data, row)
	}
	w.WriteAll(data)

	return nil
}