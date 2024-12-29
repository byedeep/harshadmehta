package types

type Transactions struct {
	Date        string
	Description string
	Enteries    []AccountEntery
}

type AccountEntery struct {
	Name   string
	Amount int
}
