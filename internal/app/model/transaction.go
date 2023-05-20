package model

import "time"

type Transaction struct {
	ID                    int       `db:"id"`
	UserID                int       `db:"user_id"`
	TransactionCategoryID int       `db:"transaction_category_id"`
	TransactionTypeID     int       `db:"transaction_type_id"`
	CurrencyID            int       `db:"currency_id"`
	TransactionDate       time.Time `db:"transaction_date"`
	Notes                 string    `db:"notes"`
	Amount                int       `db:"amount"`
}

type BrowseTransaction struct {
	UserID int
	Page   int
	Limit  int
}
