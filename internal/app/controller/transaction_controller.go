package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/meziaris/gofinance/internal/app/schema"
	"github.com/meziaris/gofinance/internal/pkg/handler"
)

type TransactionService interface {
	Create(req schema.TransactionReq) error
	BrowseAll(userID int, req schema.BrowseTransactionReq) ([]schema.GetTransactionResp, error)
	UpdateByID(id string, req schema.TransactionReq) error
	GetByID(id string) (schema.GetTransactionResp, error)
	DeleteByID(id string) error
}

type Transactionontroller struct {
	service TransactionService
}

func NewTransactionController(trxService TransactionService) Transactionontroller {
	return Transactionontroller{service: trxService}
}

func (c Transactionontroller) Create(w http.ResponseWriter, r *http.Request) {
	req := schema.TransactionReq{}

	if handler.BindAndCheck(w, r, &req) {
		return
	}

	if err := c.service.Create(req); err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp := schema.TransactionResp(req)

	handler.ResponseSuccess(w, http.StatusOK, "success create transaction type", resp)
}

func (c Transactionontroller) BrowseAll(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	userID := r.Context().Value("user_id").(string)
	id, _ := strconv.Atoi(userID)

	req := schema.BrowseTransactionReq{
		Page:  page,
		Limit: limit,
	}

	resp, err := c.service.BrowseAll(id, req)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success get transaction types", resp)
}

func (c Transactionontroller) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	req := schema.TransactionReq{}

	if handler.BindAndCheck(w, r, &req) {
		return
	}

	err := c.service.UpdateByID(id, req)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp := schema.TransactionResp(req)

	handler.ResponseSuccess(w, http.StatusOK, "success update transaction type", resp)
}

func (c Transactionontroller) Detail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resp, err := c.service.GetByID(id)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success get detail transaction type", resp)
}

func (c Transactionontroller) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := c.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success delete transaction type", nil)
}
