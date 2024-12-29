package services

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/byedeep/harshadmehta/internal/types"
)

func LoadTransactions(filename string) ([]types.Transactions, error) {
	var transaction []types.Transactions

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

	}
}
