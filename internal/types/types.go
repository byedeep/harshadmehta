package types

type Transactions struct {
	Id          int
	Date        string
	Description string
	Entries     []AccountEntry
}

type AccountEntry struct {
	Name   string
	Amount int
}
