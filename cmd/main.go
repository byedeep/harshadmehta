package types

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var input int
	var transaction Transactions
	var Entery AccountEntery
	CurrentTime := time.Now()
	CurrentDate := CurrentTime.Format("2006/01/02")

	fmt.Println("1.Create a transaction")
	fmt.Println("2.Update a transaction")
	fmt.Println("3.Delete a transaction")
	fmt.Println("4.Report a transaction")
	fmt.Println("")
	fmt.Println("Please select an option to contiune")

	fmt.Scan(&input)

	file, err := os.Open("Transaction.CSV")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	switch input {
	case 1:
		transaction.Date = CurrentDate
		fmt.Println("Please provide the discription for the transaction")
		fmt.Scan(&transaction.Description)
		fmt.Println("Please enter the name")
		fmt.Scan(&Entery.Name)
		fmt.Println("Please enter the amount")
		fmt.Scan(&Entery.Amount)

		fmt.Print(transaction)

	}

}
