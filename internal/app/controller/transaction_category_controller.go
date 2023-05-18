package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/meziaris/gofinance/internal/app/schema"
	"github.com/meziaris/gofinance/internal/pkg/handler"
)

type TransactionCategoryService interface {
	Create(category schema.TransactionCategoryReq) error
	BrowseAll(req schema.BrowseTransactionCategoryReq) ([]schema.GetTransactionCategoryResp, error)
	UpdateByID(id string, category schema.TransactionCategoryReq) error
	GetByID(id string) (schema.GetTransactionCategoryResp, error)
	DeleteByID(id string) error
}

type TransactionCategoryController struct {
	service TransactionCategoryService
}

func NewTransactionCategoryController(categoryService TransactionCategoryService) TransactionCategoryController {
	return TransactionCategoryController{service: categoryService}
}

func (c TransactionCategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	req := schema.TransactionCategoryReq{}

	if handler.BindAndCheck(w, r, &req) {
		return
	}

	if err := c.service.Create(req); err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp := schema.TransactionCategoryResp(req)

	handler.ResponseSuccess(w, http.StatusOK, "success register", resp)
}

func (c TransactionCategoryController) BrowseCategory(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	req := schema.BrowseTransactionCategoryReq{
		Page:  page,
		Limit: limit,
	}

	resp, err := c.service.BrowseAll(req)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success browse category", resp)
}

func (c TransactionCategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	req := schema.TransactionCategoryReq{}

	if handler.BindAndCheck(w, r, &req) {
		return
	}

	err := c.service.UpdateByID(id, req)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp := schema.TransactionCategoryResp(req)

	handler.ResponseSuccess(w, http.StatusOK, "success register", resp)
}

func (c TransactionCategoryController) DetailCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	catResp, err := c.service.GetByID(id)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success get detail category", catResp)
}

func (c TransactionCategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := c.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success delete category", nil)
}
