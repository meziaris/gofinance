package model

type TransactionType struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type BrowseTransactionType struct {
	Page  int
	Limit int
}
