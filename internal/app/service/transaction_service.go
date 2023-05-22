package service

import (
	"errors"
	"time"

	"github.com/meziaris/gofinance/internal/app/model"
	"github.com/meziaris/gofinance/internal/app/schema"
)

type TransactionRepository interface {
	Create(trx model.Transaction) error
	GetByID(id string) (model.Transaction, error)
	Browse(search model.BrowseTransaction) ([]model.Transaction, error)
	BrowseByType(search model.BrowseTransaction) ([]model.Transaction, error)
	UpdateByID(transaction model.Transaction) error
	DeleteByID(id string) error
}

type TransactionService struct {
	trxRepo TransactionRepository
}

func NewTransactionService(trxRepo TransactionRepository) TransactionService {
	return TransactionService{trxRepo}
}

func (s TransactionService) Create(req schema.TransactionReq) error {
	trxDate, _ := time.Parse("2006-01-02", req.TransactionDate)
	insertData := model.Transaction{
		UserID:                req.UserID,
		TransactionCategoryID: req.TransactionCategoryID,
		TransactionTypeID:     req.TransactionTypeID,
		CurrencyID:            req.CurrencyID,
		TransactionDate:       trxDate,
		Notes:                 req.Notes,
		Amount:                req.Amount,
	}

	err := s.trxRepo.Create(insertData)
	if err != nil {
		return errors.New("failed to create transaction")
	}

	return nil
}

func (s TransactionService) BrowseAll(userID int, req schema.BrowseTransactionReq) ([]schema.GetTransactionResp, error) {
	var resp []schema.GetTransactionResp

	dbSearch := model.BrowseTransaction{UserID: userID, Page: req.Page, Limit: req.Limit}
	trxTypes, err := s.trxRepo.Browse(dbSearch)
	if err != nil {
		return nil, errors.New("cannot get transactions")
	}

	for _, value := range trxTypes {
		respData := schema.GetTransactionResp{
			ID:                    value.ID,
			UserID:                value.UserID,
			TransactionCategoryID: value.TransactionCategoryID,
			TransactionTypeID:     value.TransactionTypeID,
			CurrencyID:            value.CurrencyID,
			TransactionDate:       value.TransactionDate.Format("2006-01-02"),
			Notes:                 value.Notes,
			Amount:                value.Amount,
		}
		resp = append(resp, respData)
	}

	return resp, nil
}

func (s TransactionService) BrowseByType(userID int, typeID int, req schema.BrowseTransactionReq) ([]schema.GetTransactionResp, error) {
	var resp []schema.GetTransactionResp

	dbSearch := model.BrowseTransaction{UserID: userID, TypeID: typeID, Page: req.Page, Limit: req.Limit}
	trxTypes, err := s.trxRepo.BrowseByType(dbSearch)
	if err != nil {
		return nil, errors.New("cannot get transactions")
	}

	for _, value := range trxTypes {
		respData := schema.GetTransactionResp{
			ID:                    value.ID,
			UserID:                value.UserID,
			TransactionCategoryID: value.TransactionCategoryID,
			TransactionTypeID:     value.TransactionTypeID,
			CurrencyID:            value.CurrencyID,
			TransactionDate:       value.TransactionDate.Format("2006-01-02"),
			Notes:                 value.Notes,
			Amount:                value.Amount,
		}
		resp = append(resp, respData)
	}

	return resp, nil
}

func (s TransactionService) UpdateByID(id string, req schema.TransactionReq) error {
	trxDate, _ := time.Parse("2006-01-02", req.TransactionDate)
	oldData, err := s.trxRepo.GetByID(id)
	if err != nil {
		return errors.New("transaction not exist")
	}

	updateData := model.Transaction{
		ID:                    oldData.ID,
		UserID:                req.UserID,
		TransactionCategoryID: req.TransactionCategoryID,
		TransactionTypeID:     req.TransactionTypeID,
		CurrencyID:            req.CurrencyID,
		TransactionDate:       trxDate,
		Notes:                 req.Notes,
		Amount:                req.Amount,
	}

	if err := s.trxRepo.UpdateByID(updateData); err != nil {
		return errors.New("cannot update transaction")
	}

	return nil
}

func (s TransactionService) GetByID(id string) (schema.GetTransactionResp, error) {
	resp := schema.GetTransactionResp{}
	cat, err := s.trxRepo.GetByID(id)
	if err != nil {
		return resp, errors.New("transaction not exist")
	}

	resp.ID = cat.ID
	resp.UserID = cat.UserID
	resp.TransactionCategoryID = cat.TransactionCategoryID
	resp.TransactionTypeID = cat.TransactionTypeID
	resp.CurrencyID = cat.CurrencyID
	resp.TransactionDate = cat.TransactionDate.Format("2006-01-02")
	resp.Notes = cat.Notes
	resp.Amount = cat.Amount

	return resp, nil
}

func (s TransactionService) DeleteByID(id string) error {
	_, err := s.trxRepo.GetByID(id)
	if err != nil {
		return errors.New("transaction type not exist")
	}

	if err := s.trxRepo.DeleteByID(id); err != nil {
		return errors.New("cannot delete transaction")
	}

	return nil
}
