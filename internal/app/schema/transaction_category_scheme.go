package schema

type TransactionCategoryReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type TransactionCategoryResp struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetTransactionCategoryResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BrowseTransactionCategoryReq struct {
	Page  int
	Limit int
}
