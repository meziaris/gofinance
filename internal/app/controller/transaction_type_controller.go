package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/meziaris/gofinance/internal/app/schema"
	"github.com/meziaris/gofinance/internal/pkg/handler"
)

type TransactionTypeService interface {
	Create(req schema.TransactionTypeReq) error
	BrowseAll(req schema.BrowseTransactionTypeReq) ([]schema.GetTransactionTypeResp, error)
	UpdateByID(id string, req schema.TransactionTypeReq) error
	GetByID(id string) (schema.GetTransactionTypeResp, error)
	DeleteByID(id string) error
}

type TransactionTypeController struct {
	service TransactionTypeService
}

func NewTransactionTypeController(trxTypeService TransactionTypeService) TransactionTypeController {
	return TransactionTypeController{service: trxTypeService}
}

func (c TransactionTypeController) Create(w http.ResponseWriter, r *http.Request) {
	req := schema.TransactionTypeReq{}

	if handler.BindAndCheck(w, r, &req) {
		return
	}

	if err := c.service.Create(req); err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp := schema.TransactionTypeResp(req)

	handler.ResponseSuccess(w, http.StatusOK, "success create transaction type", resp)
}

func (c TransactionTypeController) BrowseAll(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	req := schema.BrowseTransactionTypeReq{
		Page:  page,
		Limit: limit,
	}

	resp, err := c.service.BrowseAll(req)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success get transaction types", resp)
}

func (c TransactionTypeController) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	req := schema.TransactionTypeReq{}

	if handler.BindAndCheck(w, r, &req) {
		return
	}

	err := c.service.UpdateByID(id, req)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp := schema.TransactionTypeResp(req)

	handler.ResponseSuccess(w, http.StatusOK, "success update transaction type", resp)
}

func (c TransactionTypeController) Detail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	catResp, err := c.service.GetByID(id)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success get detail transaction type", catResp)
}

func (c TransactionTypeController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := c.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success delete transaction type", nil)
}
