package service

import (
	"errors"

	"github.com/meziaris/gofinance/internal/app/model"
	"github.com/meziaris/gofinance/internal/app/schema"
)

type CurrencyRepository interface {
	Create(currency model.Currency) error
	GetByID(id string) (model.Currency, error)
	Browse(search model.BrowseCurrency) ([]model.Currency, error)
	UpdateByID(currency model.Currency) error
	DeleteByID(id string) error
}

type CurrencyService struct {
	currencyRepo CurrencyRepository
}

func NewCurrencyService(currencyRepo CurrencyRepository) CurrencyService {
	return CurrencyService{currencyRepo: currencyRepo}
}

func (s CurrencyService) Create(req schema.CurrencyReq) error {
	insertData := model.Currency{Name: req.Name}

	err := s.currencyRepo.Create(insertData)
	if err != nil {
		return errors.New("failed to create currency")
	}

	return nil
}

func (s CurrencyService) BrowseAll(req schema.BrowseCurrencyReq) ([]schema.GetCurrencyResp, error) {
	var resp []schema.GetCurrencyResp

	dbSearch := model.BrowseCurrency{Page: req.Page, Limit: req.Limit}
	categories, err := s.currencyRepo.Browse(dbSearch)
	if err != nil {
		return nil, errors.New("cannot get categories")
	}

	for _, value := range categories {
		var respData schema.GetCurrencyResp
		respData.ID = value.ID
		respData.Name = value.Name
		resp = append(resp, respData)
	}

	return resp, nil
}

func (s CurrencyService) UpdateByID(id string, req schema.CurrencyReq) error {
	oldData, err := s.currencyRepo.GetByID(id)
	if err != nil {
		return errors.New("currency not exist")
	}

	updateDate := model.Currency{
		ID:   oldData.ID,
		Name: req.Name,
	}

	if err := s.currencyRepo.UpdateByID(updateDate); err != nil {
		return errors.New("cannot update currency")
	}

	return nil
}

func (s CurrencyService) GetByID(id string) (schema.GetCurrencyResp, error) {
	resp := schema.GetCurrencyResp{}
	cat, err := s.currencyRepo.GetByID(id)
	if err != nil {
		return resp, errors.New("currency not exist")
	}

	resp.ID = cat.ID
	resp.Name = cat.Name

	return resp, nil
}

func (s CurrencyService) DeleteByID(id string) error {
	_, err := s.currencyRepo.GetByID(id)
	if err != nil {
		return errors.New("currency not exist")
	}

	if err := s.currencyRepo.DeleteByID(id); err != nil {
		return errors.New("cannot delete currency")
	}

	return nil
}
