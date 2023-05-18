package service

import (
	"errors"

	"github.com/meziaris/gofinance/internal/app/model"
	"github.com/meziaris/gofinance/internal/app/schema"
)

type TransactionCategoryRepository interface {
	Create(category model.TransactionCategory) error
	Browse(search model.BrowseTransactionCategory) ([]model.TransactionCategory, error)
	GetByID(id string) (model.TransactionCategory, error)
	UpdateByID(category model.TransactionCategory) error
	DeleteByID(id string) error
}

type TransactionCategoryService struct {
	categoryRepo TransactionCategoryRepository
}

func NewTransactionCategoryService(categoryRepo TransactionCategoryRepository) TransactionCategoryService {
	return TransactionCategoryService{categoryRepo: categoryRepo}
}

func (s TransactionCategoryService) Create(req schema.TransactionCategoryReq) error {
	insertData := model.TransactionCategory{
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.categoryRepo.Create(insertData)
	if err != nil {
		return err
	}

	return nil
}

func (s TransactionCategoryService) BrowseAll(req schema.BrowseTransactionCategoryReq) ([]schema.GetTransactionCategoryResp, error) {
	var resp []schema.GetTransactionCategoryResp

	dbSearch := model.BrowseTransactionCategory{Page: req.Page, Limit: req.Limit}
	categories, err := s.categoryRepo.Browse(dbSearch)
	if err != nil {
		return nil, errors.New("cannot get categories")
	}

	for _, value := range categories {
		var respData schema.GetTransactionCategoryResp
		respData.ID = value.ID
		respData.Name = value.Name
		respData.Description = value.Description
		resp = append(resp, respData)
	}

	return resp, nil
}

func (s TransactionCategoryService) Update(req schema.TransactionCategoryReq) error {
	insertData := model.TransactionCategory{
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.categoryRepo.Create(insertData)
	if err != nil {
		return err
	}

	return nil
}

func (s TransactionCategoryService) UpdateByID(id string, req schema.TransactionCategoryReq) error {
	oldData, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return errors.New("category not exist")
	}

	updateDate := model.TransactionCategory{
		ID:          oldData.ID,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.categoryRepo.UpdateByID(updateDate); err != nil {
		return errors.New("cannot update category")
	}

	return nil
}

func (s TransactionCategoryService) GetByID(id string) (schema.GetTransactionCategoryResp, error) {
	resp := schema.GetTransactionCategoryResp{}
	cat, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return resp, errors.New("category not exist")
	}

	resp.ID = cat.ID
	resp.Name = cat.Name
	resp.Description = cat.Description

	return resp, nil
}

func (s TransactionCategoryService) DeleteByID(id string) error {
	_, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return errors.New("category not exist")
	}

	if err := s.categoryRepo.DeleteByID(id); err != nil {
		return errors.New("cannot delete category")
	}

	return nil
}
