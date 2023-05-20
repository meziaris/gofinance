package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meziaris/gofinance/internal/app/model"
	log "github.com/sirupsen/logrus"
)

type TransactionRepository struct {
	DB *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return TransactionRepository{DB: db}
}

func (r TransactionRepository) Create(trx model.Transaction) error {
	var (
		sqlStatement = `
			INSERT INTO transactions(
				user_id,
				transaction_category_id,
				transaction_type_id,
				currency_id,
				transaction_date,
				notes,
				amount
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`
	)

	_, err := r.DB.Exec(
		sqlStatement,
		trx.UserID,
		trx.TransactionCategoryID,
		trx.TransactionTypeID,
		trx.CurrencyID,
		trx.TransactionDate,
		trx.Notes,
		trx.Amount,
	)
	if err != nil {
		log.Error("error TransactionRepository - Create : %w", err)
		return err
	}

	return nil
}

func (r TransactionRepository) GetByID(id string) (model.Transaction, error) {
	var (
		trx          = model.Transaction{}
		sqlStatement = `
			SELECT id, user_id, transaction_category_id, transaction_type_id, currency_id, transaction_date, notes, amount
			FROM transactions
			WHERE id = $1
			LIMIT 1
		`
	)

	if err := r.DB.QueryRowx(sqlStatement, id).StructScan(&trx); err != nil {
		log.Error("error TransactionRepository - GetByID :", err)
		return trx, err
	}

	return trx, nil
}

func (r TransactionRepository) Browse(search model.BrowseTransaction) ([]model.Transaction, error) {
	var (
		limit        = search.Limit
		offset       = limit * (search.Page - 1)
		transactions []model.Transaction
		sqlStatement = `
			SELECT
				trx.id,
				trx.user_id,
				trx.transaction_category_id,
				trx.transaction_type_id,
				trx.currency_id,
				trx.transaction_date,
				trx.notes,
				trx.amount
			FROM transactions trx
			LEFT JOIN transaction_types as trxtypes
				ON trx.transaction_type_id = trxtypes.id
			LEFT JOIN transaction_categories as trxcat
				ON trx.transaction_category_id = trxcat.id
			LEFT JOIN currencies
				ON trx.currency_id = currencies.id
			WHERE trx.user_id = $1
			LIMIT $2
			OFFSET $3
		`
	)

	rows, err := r.DB.Queryx(sqlStatement, search.UserID, limit, offset)
	if err != nil {
		log.Error(fmt.Errorf("error TransactionRepository - Browse : %w", err))
		return transactions, err
	}

	for rows.Next() {
		var transaction model.Transaction
		err := rows.StructScan(&transaction)
		if err != nil {
			log.Error(fmt.Errorf("error TransactionRepository - Browse : %w", err))
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r TransactionRepository) UpdateByID(transaction model.Transaction) error {
	var (
		sqlStatement = `
			UPDATE transactions
			SET updated_at = NOW(),
				user_id = $2,
				transaction_category_id = $3,
				transaction_type_id = $4,
				currency_id = $5,
				transaction_date = $6,
				notes = $7,
				amount = $8
			WHERE id = $1
		`
	)

	result, err := r.DB.Exec(
		sqlStatement,
		transaction.ID,
		transaction.UserID,
		transaction.TransactionCategoryID,
		transaction.TransactionTypeID,
		transaction.CurrencyID,
		transaction.TransactionDate,
		transaction.Notes,
		transaction.Amount,
	)
	if err != nil {
		log.Error(fmt.Errorf("error TransactionRepository - UpdateByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}

func (r TransactionRepository) DeleteByID(id string) error {
	var (
		sqlStatement = `
			DELETE FROM transactions
			WHERE id = $1
		`
	)

	result, err := r.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error TransactionRepository - DeleteByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}
