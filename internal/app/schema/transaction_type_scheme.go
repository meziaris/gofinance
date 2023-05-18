package schema

type TransactionTypeReq struct {
	Name string `json:"name" validate:"required"`
}

type TransactionTypeResp struct {
	Name string `json:"name"`
}

type GetTransactionTypeResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type BrowseTransactionTypeReq struct {
	Page  int
	Limit int
}
