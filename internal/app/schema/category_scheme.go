package schema

type CategoryReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CategoryResp struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetCategoryResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BrowseCategoryReq struct {
	Page  int
	Limit int
}
