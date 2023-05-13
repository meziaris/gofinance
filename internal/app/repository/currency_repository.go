package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meziaris/gofinance/internal/app/model"
	log "github.com/sirupsen/logrus"
)

type CurrencyRepository struct {
	DB *sqlx.DB
}

func NewCurrencyRepository(db *sqlx.DB) CurrencyRepository {
	return CurrencyRepository{DB: db}
}

func (r CurrencyRepository) Create(currency model.Currency) error {
	var (
		sqlStatement = `
			INSERT INTO currencies(name)
			VALUES ($1)
		`
	)

	_, err := r.DB.Exec(sqlStatement, currency.Name)
	if err != nil {
		log.Error("error CurrencyRepository - Create : %w", err)
		return err
	}

	return nil
}

func (r CurrencyRepository) GetByID(id string) (model.Currency, error) {
	var (
		currency     = model.Currency{}
		sqlStatement = `
			SELECT id, name
			FROM currencies
			WHERE id = $1
			LIMIT 1
		`
	)

	if err := r.DB.QueryRowx(sqlStatement, id).StructScan(&currency); err != nil {
		log.Error("error CurrencyRepository - GetByID :", err)
		return currency, err
	}

	return currency, nil
}

func (cr CurrencyRepository) Browse(search model.BrowseCurrency) ([]model.Currency, error) {
	var (
		limit        = search.Limit
		offset       = limit * (search.Page - 1)
		currencies   []model.Currency
		sqlStatement = `
			SELECT id, name
			FROM currencies
			LIMIT $1
			OFFSET $2
		`
	)

	rows, err := cr.DB.Queryx(sqlStatement, limit, offset)
	if err != nil {
		log.Error(fmt.Errorf("error CurrencyRepository - Browse : %w", err))
		return currencies, err
	}

	for rows.Next() {
		var currency model.Currency
		err := rows.StructScan(&currency)
		if err != nil {
			log.Error(fmt.Errorf("error CurrencyRepository - Browse : %w", err))
		}
		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (r CurrencyRepository) UpdateByID(currency model.Currency) error {
	var (
		sqlStatement = `
			UPDATE currencies
			SET updated_at = NOW(),
				name = $2
			WHERE id = $1
		`
	)

	result, err := r.DB.Exec(sqlStatement, currency.ID, currency.Name)
	if err != nil {
		log.Error(fmt.Errorf("error CurrencyRepository - UpdateByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}

func (r CurrencyRepository) DeleteByID(id string) error {
	var (
		sqlStatement = `
			DELETE FROM currencies
			WHERE id = $1
		`
	)

	result, err := r.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error CurrencyRepository - DeleteByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}
