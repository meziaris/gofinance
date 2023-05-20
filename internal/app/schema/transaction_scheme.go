package schema

type TransactionReq struct {
	UserID                int    `json:"user_id" validate:"required"`
	TransactionCategoryID int    `json:"transaction_category_id" validate:"required"`
	TransactionTypeID     int    `json:"transaction_type_id" validate:"required"`
	CurrencyID            int    `json:"currency_id" validate:"required"`
	TransactionDate       string `json:"transaction_date" validate:"required"`
	Notes                 string `json:"notes" validate:"required"`
	Amount                int    `json:"amount" validate:"required,number"`
}

type TransactionResp struct {
	UserID                int    `json:"user_id"`
	TransactionCategoryID int    `json:"transaction_category_id"`
	TransactionTypeID     int    `json:"transaction_type_id"`
	CurrencyID            int    `json:"currency_id"`
	TransactionDate       string `json:"transaction_date"`
	Notes                 string `json:"notes"`
	Amount                int    `json:"amount"`
}

type GetTransactionResp struct {
	ID                    int    `json:"id"`
	UserID                int    `json:"user_id"`
	TransactionCategoryID int    `json:"transaction_category_id"`
	TransactionTypeID     int    `json:"transaction_type_id"`
	CurrencyID            int    `json:"currency_id"`
	TransactionDate       string `json:"transaction_date"`
	Notes                 string `json:"notes"`
	Amount                int    `json:"amount"`
}

type BrowseTransactionReq struct {
	Page  int
	Limit int
}
