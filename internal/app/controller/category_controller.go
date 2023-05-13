package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/meziaris/gofinance/internal/app/schema"
	"github.com/meziaris/gofinance/internal/pkg/handler"
)

type CategoryService interface {
	Create(category schema.CategoryReq) error
	BrowseAll(req schema.BrowseCategoryReq) ([]schema.GetCategoryResp, error)
	UpdateByID(id string, category schema.CategoryReq) error
	GetByID(id string) (schema.GetCategoryResp, error)
	DeleteByID(id string) error
}

type CategoryController struct {
	service CategoryService
}

func NewCategoryController(categoryService CategoryService) CategoryController {
	return CategoryController{service: categoryService}
}

func (c CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	req := schema.CategoryReq{}

	if handler.BindAndCheck(w, r, &req) {
		return
	}

	if err := c.service.Create(req); err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp := schema.CategoryResp(req)

	handler.ResponseSuccess(w, http.StatusOK, "success register", resp)
}

func (c CategoryController) BrowseCategory(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	req := schema.BrowseCategoryReq{
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

func (c CategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	req := schema.CategoryReq{}

	if handler.BindAndCheck(w, r, &req) {
		return
	}

	err := c.service.UpdateByID(id, req)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp := schema.CategoryResp(req)

	handler.ResponseSuccess(w, http.StatusOK, "success register", resp)
}

func (c CategoryController) DetailCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	catResp, err := c.service.GetByID(id)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success get detail category", catResp)
}

func (c CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := c.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success delete category", nil)
}
