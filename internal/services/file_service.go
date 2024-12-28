package services

import (
	"encoding/csv"
	"os"
	"strconv"
)

func LoadTransactions(filename string) ([]Transactions, error) {
	var transaction []Transactions

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
	var currentTransaction Transactions

	for _, record := range records {
		if len(record) < 4 {
			continue
		}
	}

	id, _ := strconv.Atoi(record[0])
	if err != nil {
		contiune
	}

}
