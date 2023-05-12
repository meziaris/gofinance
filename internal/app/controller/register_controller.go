package controller

import (
	"net/http"

	"github.com/meziaris/gofinance/internal/app/schema"
	"github.com/meziaris/gofinance/internal/pkg/handler"
)

type Registerer interface {
	Register(req schema.RegisterReq) error
}

type RegistrationController struct {
	service Registerer
}

func NewRegistrationController(userService Registerer) RegistrationController {
	return RegistrationController{service: userService}
}

func (c RegistrationController) Register(w http.ResponseWriter, r *http.Request) {
	req := schema.RegisterReq{}

	if handler.BindAndCheck(w, r, &req) {
		return
	}

	err := c.service.Register(req)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp := schema.RegisterResp{
		Username: req.Username,
		FullName: req.FullName,
	}

	handler.ResponseSuccess(w, http.StatusOK, "success register", resp)
}
