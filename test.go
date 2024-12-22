package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type AccountEntry struct {
	Name   string
	Amount int
}

type Transaction struct {
	Id          int
	Date        string
	Description string
	Entries     []AccountEntry
}

func loadTransactions(fileName string) ([]Transaction, error) {
	var transactions []Transaction

	file, err := os.Open(fileName)
	if err != nil {
		// If the file doesn't exist, return an empty slice
		if os.IsNotExist(err) {
			return transactions, nil
		}
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	currentId := -1
	var currentTransaction Transaction

	for _, record := range records {
		if len(record) < 4 {
			continue // Skip malformed rows
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}

		if id != currentId {
			// Save the previous transaction if it's complete
			if currentTransaction.Id != 0 {
				transactions = append(transactions, currentTransaction)
			}

			// Start a new transaction
			currentId = id
			currentTransaction = Transaction{
				Id:          id,
				Date:        record[1],
				Description: record[2],
				Entries:     []AccountEntry{},
			}
		}

		amount, err := strconv.Atoi(record[4])
		if err != nil {
			continue
		}

		currentTransaction.Entries = append(currentTransaction.Entries, AccountEntry{
			Name:   record[3],
			Amount: amount,
		})
	}

	// Append the last transaction
	if currentTransaction.Id != 0 {
		transactions = append(transactions, currentTransaction)
	}

	return transactions, nil
}

func saveTransactions(fileName string, transactions []Transaction) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, transaction := range transactions {
		for _, entry := range transaction.Entries {
			record := []string{
				strconv.Itoa(transaction.Id),
				transaction.Date,
				transaction.Description,
				entry.Name,
				strconv.Itoa(entry.Amount),
			}
			if err := writer.Write(record); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	fileName := "Transaction.csv"
	var transactions []Transaction
	var err error

	transactions, err = loadTransactions(fileName)
	if err != nil {
		fmt.Println("Error loading transactions:", err)
		return
	}

	fmt.Println("1. Create a transaction")
	fmt.Println("2. Update a transaction")
	fmt.Println("3. Delete a transaction")
	fmt.Println("4. View all transactions")
	fmt.Println("")
	fmt.Println("Please select an option to continue:")

	var input int
	fmt.Scan(&input)

	switch input {
	case 1:
		// Create a new transaction
		var newTransaction Transaction
		CurrentTime := time.Now()
		newTransaction.Date = CurrentTime.Format("2006/01/02")

		fmt.Println("Enter a description for the transaction:")
		fmt.Scan(&newTransaction.Description)

		newTransaction.Id = len(transactions) + 1 // Assign a new ID

		for {
			var entry AccountEntry
			fmt.Println("Enter the account name (or type 'done' to finish):")
			fmt.Scan(&entry.Name)
			if entry.Name == "done" {
				break
			}

			fmt.Println("Enter the amount:")
			fmt.Scan(&entry.Amount)
			newTransaction.Entries = append(newTransaction.Entries, entry)
		}

		transactions = append(transactions, newTransaction)
		fmt.Println("Transaction created successfully!")

	case 2:
		// Update an existing transaction
		var id int
		fmt.Println("Enter the transaction ID to update:")
		fmt.Scan(&id)

		for i, transaction := range transactions {
			if transaction.Id == id {
				fmt.Println("Updating transaction:", transaction)
				fmt.Println("Enter a new description (or leave empty to keep current):")
				var newDescription string
				fmt.Scan(&newDescription)
				if newDescription != "" {
					transactions[i].Description = newDescription
				}

				fmt.Println("Do you want to edit entries? (yes/no):")
				var editEntries string
				fmt.Scan(&editEntries)
				if editEntries == "yes" {
					var entries []AccountEntry
					for {
						var entry AccountEntry
						fmt.Println("Enter the account name (or type 'done' to finish):")
						fmt.Scan(&entry.Name)
						if entry.Name == "done" {
							break
						}

						fmt.Println("Enter the amount:")
						fmt.Scan(&entry.Amount)
						entries = append(entries, entry)
					}
					transactions[i].Entries = entries
				}
				fmt.Println("Transaction updated successfully!")
				break
			}
		}

	case 3:
		// Delete a transaction
		var id int
		fmt.Println("Enter the transaction ID to delete:")
		fmt.Scan(&id)

		for i, transaction := range transactions {
			if transaction.Id == id {
				transactions = append(transactions[:i], transactions[i+1:]...)
				fmt.Println("Transaction deleted successfully!")
				break
			}
		}

	case 4:
		// View all transactions
		fmt.Println("All Transactions:")
		for _, transaction := range transactions {
			fmt.Printf("ID: %d, Date: %s, Description: %s\n", transaction.Id, transaction.Date, transaction.Description)
			for _, entry := range transaction.Entries {
				fmt.Printf("    Account: %s, Amount: %d\n", entry.Name, entry.Amount)
			}
		}
	}

	// Save transactions back to the CSV file
	err = saveTransactions(fileName, transactions)
	if err != nil {
		fmt.Println("Error saving transactions:", err)
	} else {
		fmt.Println("Transactions saved successfully!")
	}
}
