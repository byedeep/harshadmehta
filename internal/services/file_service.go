package services

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/byedeep/harshadmehta/internal/types"
)

func LoadTransactions(filename string) ([]types.Transactions, error) {
	var transactions []types.Transactions
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	CurrentId := -1
	var currentTransaction types.Transactions

	for _, record := range records {
		if len(record) < 4 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}

		if id != CurrentId {
			if currentTransaction.Id != 0 {
				transactions = append(transactions, currentTransaction)
			}

			CurrentId = id
			currentTransaction = types.Transactions{

				Id:          id,
				Date:        record[1],
				Description: record[2],
				Entries:     []types.AccountEntry{},
			}
		}
		amount, err := strconv.Atoi(record[4])
		if err != nil {
			return nil, err
		}
		currentTransaction.Entries = append(currentTransaction.Entries, types.AccountEntry{
			Name:   record[3],
			Amount: amount,
		})
		if currentTransaction.Id != 0 {
			transactions = append(transactions, currentTransaction)
		}

	}
	return transactions, nil
}

func SaveTransaction(filename string, transactions []types.Transactions) error {
	file, err := os.Create(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, transaction := range transactions {
		for _, entery := range transaction.Entries {
			record := []string{
				strconv.Itoa(transaction.Id),
				transaction.Date,
				transaction.Description,
				entery.Name,
				strconv.Itoa(entery.Amount),
			}
			if err := writer.Write(record); err != nil {
				return nil
			}

		}
	}
	return nil
}
