package schema

type CurrencyReq struct {
	Name string `json:"name" validate:"required"`
}

type CurrencyResp struct {
	Name string `json:"name"`
}

type GetCurrencyResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type BrowseCurrencyReq struct {
	Page  int
	Limit int
}
