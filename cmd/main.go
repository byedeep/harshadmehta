package main

import (
	"fmt"
	"time"

	"github.com/byedeep/harshadmehta/internal/services"
	"github.com/byedeep/harshadmehta/internal/types"
)

func main() {
	var input int
	var Transactions []types.Transactions
	var Entry types.AccountEntry
	filename := "../data/transactions.csv"
	Transactions, err := services.LoadTransactions(filename)
	if err != nil {
		fmt.Print("error", err)
		return
	}

	CurrentTime := time.Now()
	CurrentDate := CurrentTime.Format("2006/01/02")

	fmt.Println("1.Create a transaction")
	fmt.Println("2.Update a transaction")
	fmt.Println("3.Delete a transaction")
	fmt.Println("4.Report all transaction")
	fmt.Println("")
	fmt.Println("Please select an option to contiune")

	fmt.Scan(&input)
	switch input {
	case 1:
		var NewTransaction types.Transactions

		NewTransaction.Date = CurrentDate
		NewTransaction.Id = len(Transactions) + 1
		fmt.Println("Please enter the Description:")
		fmt.Scan(&NewTransaction.Description)

		for {
			fmt.Println("Please enter the name or type done to finish")
			fmt.Scan(&Entry.Name)
			if Entry.Name == "done" {
				break
			}
			fmt.Println("Please enter the amount")
			fmt.Scan(&Entry.Amount)
			NewTransaction.Entries = append(NewTransaction.Entries, Entry)
		}

		Transactions = append(Transactions, NewTransaction)
		fmt.Println("Transaction created!")
	case 2:
		var id int

		fmt.Println("Please enter the id:")
		fmt.Scan(&id)

		for i, transaction := range Transactions {
			if transaction.Id == id {
				fmt.Println("Updating transaction:", transaction)
				fmt.Println("Enter the new Description:")
				var NewDescription string
				fmt.Scan(&NewDescription)
				if NewDescription != "" {
					Transactions[i].Description = NewDescription
				}

				fmt.Println("Do you want to edit entries?[y/n]:")
				var editEntries string
				fmt.Scan(&editEntries)
				if editEntries == "y" {
					var entries []types.AccountEntry
					for {
						var entry types.AccountEntry
						fmt.Println("Please enter the name or type done to finish")
						fmt.Scan(&entry.Name)
						if entry.Name == "done" {
							break
						}

						fmt.Println("Please enter the amount:")
						fmt.Scan(&entry.Amount)
						entries = append(entries, entry)
					}
					Transactions[i].Entries = entries
				}
				fmt.Println("Transaction updated!")
				break
			}
		}
	case 3:
		var id int
		fmt.Println("Enter the transaction ID to delete:")
		fmt.Scan(&id)

		for i, transaction := range Transactions {
			if transaction.Id == id {
				Transactions = append(Transactions[:i], Transactions[i+1:]...)
				fmt.Println("Transaction deleted successfully!")
				break
			}
		}
	case 4:

		fmt.Println("All Transactions:")
		for _, transaction := range Transactions {
			fmt.Printf("ID: %d, Date: %s, Description: %s\n", transaction.Id, transaction.Date, transaction.Description)
			for _, entry := range transaction.Entries {
				fmt.Printf("    Account: %s, Amount: %d\n", entry.Name, entry.Amount)
			}
		}

	}
	err = services.SaveTransaction("../data/transactions.csv", Transactions)
	if err != nil {
		fmt.Println("Error!:", err)
	} else {
		fmt.Println("Transactions saved!")
	}
}
