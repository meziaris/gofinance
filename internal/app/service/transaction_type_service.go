package service

import (
	"errors"

	"github.com/meziaris/gofinance/internal/app/model"
	"github.com/meziaris/gofinance/internal/app/schema"
)

type TransactionTypeRepository interface {
	Create(category model.TransactionType) error
	Browse(search model.BrowseTransactionType) ([]model.TransactionType, error)
	GetByID(id string) (model.TransactionType, error)
	UpdateByID(category model.TransactionType) error
	DeleteByID(id string) error
}

type TransactionTypeService struct {
	trxTypeRepo TransactionTypeRepository
}

func NewTransactionTypeService(trxTypeRepo TransactionTypeRepository) TransactionTypeService {
	return TransactionTypeService{trxTypeRepo: trxTypeRepo}
}

func (s TransactionTypeService) Create(req schema.TransactionTypeReq) error {
	insertData := model.TransactionType{
		Name: req.Name,
	}

	err := s.trxTypeRepo.Create(insertData)
	if err != nil {
		return err
	}

	return nil
}

func (s TransactionTypeService) BrowseAll(req schema.BrowseTransactionTypeReq) ([]schema.GetTransactionTypeResp, error) {
	var resp []schema.GetTransactionTypeResp

	dbSearch := model.BrowseTransactionType{Page: req.Page, Limit: req.Limit}
	trxTypes, err := s.trxTypeRepo.Browse(dbSearch)
	if err != nil {
		return nil, errors.New("cannot get transaction types")
	}

	for _, value := range trxTypes {
		var respData schema.GetTransactionTypeResp
		respData.ID = value.ID
		respData.Name = value.Name
		resp = append(resp, respData)
	}

	return resp, nil
}

func (s TransactionTypeService) UpdateByID(id string, req schema.TransactionTypeReq) error {
	oldData, err := s.trxTypeRepo.GetByID(id)
	if err != nil {
		return errors.New("transaction type not exist")
	}

	updateDate := model.TransactionType{
		ID:   oldData.ID,
		Name: req.Name,
	}

	if err := s.trxTypeRepo.UpdateByID(updateDate); err != nil {
		return errors.New("cannot update transaction type")
	}

	return nil
}

func (s TransactionTypeService) GetByID(id string) (schema.GetTransactionTypeResp, error) {
	resp := schema.GetTransactionTypeResp{}
	cat, err := s.trxTypeRepo.GetByID(id)
	if err != nil {
		return resp, errors.New("transaction type not exist")
	}

	resp.ID = cat.ID
	resp.Name = cat.Name

	return resp, nil
}

func (s TransactionTypeService) DeleteByID(id string) error {
	_, err := s.trxTypeRepo.GetByID(id)
	if err != nil {
		return errors.New("transaction type not exist")
	}

	if err := s.trxTypeRepo.DeleteByID(id); err != nil {
		return errors.New("cannot delete transaction type")
	}

	return nil
}
