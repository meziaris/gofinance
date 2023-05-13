package service

import (
	"errors"

	"github.com/meziaris/gofinance/internal/app/model"
	"github.com/meziaris/gofinance/internal/app/schema"
)

type CategoryRepository interface {
	Create(category model.Category) error
	Browse(search model.BrowseCategory) ([]model.Category, error)
	GetByID(id string) (model.Category, error)
	UpdateByID(category model.Category) error
	DeleteByID(id string) error
}

type CategoryService struct {
	categoryRepo CategoryRepository
}

func NewCategoryService(categoryRepo CategoryRepository) CategoryService {
	return CategoryService{categoryRepo: categoryRepo}
}

func (s CategoryService) Create(req schema.CategoryReq) error {
	insertData := model.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.categoryRepo.Create(insertData)
	if err != nil {
		return err
	}

	return nil
}

func (s CategoryService) BrowseAll(req schema.BrowseCategoryReq) ([]schema.GetCategoryResp, error) {
	var resp []schema.GetCategoryResp

	dbSearch := model.BrowseCategory{Page: req.Page, Limit: req.Limit}
	categories, err := s.categoryRepo.Browse(dbSearch)
	if err != nil {
		return nil, errors.New("cannot get categories")
	}

	for _, value := range categories {
		var respData schema.GetCategoryResp
		respData.ID = value.ID
		respData.Name = value.Name
		respData.Description = value.Description
		resp = append(resp, respData)
	}

	return resp, nil
}

func (s CategoryService) Update(req schema.CategoryReq) error {
	insertData := model.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.categoryRepo.Create(insertData)
	if err != nil {
		return err
	}

	return nil
}

func (s CategoryService) UpdateByID(id string, req schema.CategoryReq) error {
	oldData, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return errors.New("category not exist")
	}

	updateDate := model.Category{
		ID:          oldData.ID,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.categoryRepo.UpdateByID(updateDate); err != nil {
		return errors.New("cannot update category")
	}

	return nil
}

func (s CategoryService) GetByID(id string) (schema.GetCategoryResp, error) {
	resp := schema.GetCategoryResp{}
	cat, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return resp, errors.New("category not exist")
	}

	resp.ID = cat.ID
	resp.Name = cat.Name
	resp.Description = cat.Description

	return resp, nil
}

func (s CategoryService) DeleteByID(id string) error {
	_, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return errors.New("category not exist")
	}

	if err := s.categoryRepo.DeleteByID(id); err != nil {
		return errors.New("cannot delete category")
	}

	return nil
}
