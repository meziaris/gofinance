package model

type TransactionCategory struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type BrowseTransactionCategory struct {
	Page  int
	Limit int
}
