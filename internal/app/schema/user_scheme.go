package schema

type RegisterReq struct {
	Username string `json:"username" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterResp struct {
	Username string `json:"username" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
}
