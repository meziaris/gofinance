package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/meziaris/gofinance/internal/app/schema"
	"github.com/meziaris/gofinance/internal/pkg/handler"
)

type CurrencyService interface {
	Create(req schema.CurrencyReq) error
	BrowseAll(req schema.BrowseCurrencyReq) ([]schema.GetCurrencyResp, error)
	UpdateByID(id string, req schema.CurrencyReq) error
	GetByID(id string) (schema.GetCurrencyResp, error)
	DeleteByID(id string) error
}

type CurrencyController struct {
	service CurrencyService
}

func NewCurrencyController(currencyService CurrencyService) CurrencyController {
	return CurrencyController{service: currencyService}
}

func (c CurrencyController) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	req := schema.CurrencyReq{}

	if handler.BindAndCheck(w, r, &req) {
		return
	}

	if err := c.service.Create(req); err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp := schema.CurrencyResp(req)

	handler.ResponseSuccess(w, http.StatusOK, "success create currency", resp)
}

func (c CurrencyController) BrowseCurrency(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	req := schema.BrowseCurrencyReq{
		Page:  page,
		Limit: limit,
	}

	resp, err := c.service.BrowseAll(req)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success browse currency", resp)
}

func (c CurrencyController) UpdateCurrency(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	req := schema.CurrencyReq{}

	if handler.BindAndCheck(w, r, &req) {
		return
	}

	err := c.service.UpdateByID(id, req)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp := schema.CurrencyResp(req)

	handler.ResponseSuccess(w, http.StatusOK, "success update currency", resp)
}

func (c CurrencyController) DetailCurrency(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	catResp, err := c.service.GetByID(id)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success get detail currency", catResp)
}

func (c CurrencyController) DeleteCurrency(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := c.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success delete currency", nil)
}
