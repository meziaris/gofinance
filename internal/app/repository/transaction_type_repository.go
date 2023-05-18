package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meziaris/gofinance/internal/app/model"
	log "github.com/sirupsen/logrus"
)

type TransactionTypeRepository struct {
	DB *sqlx.DB
}

func NewTransactionTypeRepository(db *sqlx.DB) TransactionTypeRepository {
	return TransactionTypeRepository{DB: db}
}

func (r TransactionTypeRepository) Create(transactionType model.TransactionType) error {
	var (
		sqlStatement = `
			INSERT INTO transaction_types(name)
			VALUES ($1)
		`
	)

	_, err := r.DB.Exec(sqlStatement, transactionType.Name)
	if err != nil {
		log.Error("error TransactionTypeRepository - Create : %w", err)
		return err
	}

	return nil
}

func (r TransactionTypeRepository) GetByID(id string) (model.TransactionType, error) {
	var (
		transactionType = model.TransactionType{}
		sqlStatement    = `
			SELECT id, name
			FROM transaction_types
			WHERE id = $1
			LIMIT 1
		`
	)

	if err := r.DB.QueryRowx(sqlStatement, id).StructScan(&transactionType); err != nil {
		log.Error("error TransactionTypeRepository - GetByID :", err)
		return transactionType, err
	}

	return transactionType, nil
}

func (r TransactionTypeRepository) Browse(search model.BrowseTransactionType) ([]model.TransactionType, error) {
	var (
		limit            = search.Limit
		offset           = limit * (search.Page - 1)
		transactionTypes []model.TransactionType
		sqlStatement     = `
			SELECT id, name
			FROM transaction_types
			LIMIT $1
			OFFSET $2
		`
	)

	rows, err := r.DB.Queryx(sqlStatement, limit, offset)
	if err != nil {
		log.Error(fmt.Errorf("error TransactionTypeRepository - Browse : %w", err))
		return transactionTypes, err
	}

	for rows.Next() {
		var transactionType model.TransactionType
		err := rows.StructScan(&transactionType)
		if err != nil {
			log.Error(fmt.Errorf("error TransactionTypeRepository - Browse : %w", err))
		}
		transactionTypes = append(transactionTypes, transactionType)
	}

	return transactionTypes, nil
}

func (r TransactionTypeRepository) UpdateByID(transactionType model.TransactionType) error {
	var (
		sqlStatement = `
			UPDATE transaction_types
			SET updated_at = NOW(),
				name = $2
			WHERE id = $1
		`
	)

	result, err := r.DB.Exec(sqlStatement, transactionType.ID, transactionType.Name)
	if err != nil {
		log.Error(fmt.Errorf("error TransactionTypeRepository - UpdateByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}

func (r TransactionTypeRepository) DeleteByID(id string) error {
	var (
		sqlStatement = `
			DELETE FROM transaction_types
			WHERE id = $1
		`
	)

	result, err := r.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error TransactionTypeRepository - DeleteByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}
