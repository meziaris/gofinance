package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meziaris/gofinance/internal/app/model"
	log "github.com/sirupsen/logrus"
)

type TransactionCategoryRepository struct {
	DB *sqlx.DB
}

func NewTransactionCategoryRepository(db *sqlx.DB) TransactionCategoryRepository {
	return TransactionCategoryRepository{DB: db}
}

func (r TransactionCategoryRepository) Create(category model.TransactionCategory) error {
	var (
		sqlStatement = `
			INSERT INTO transaction_categories(name, description)
			VALUES ($1, $2)
		`
	)

	_, err := r.DB.Exec(sqlStatement, category.Name, category.Description)
	if err != nil {
		log.Error("error CategoryRepository - Create : %w", err)
		return err
	}

	return nil
}

func (r TransactionCategoryRepository) GetByID(id string) (model.TransactionCategory, error) {
	var (
		category     = model.TransactionCategory{}
		sqlStatement = `
			SELECT id, name, description
			FROM transaction_categories
			WHERE id = $1
			LIMIT 1
		`
	)

	if err := r.DB.QueryRowx(sqlStatement, id).StructScan(&category); err != nil {
		log.Error("error CategoryRepository - GetByID :", err)
		return category, err
	}

	return category, nil
}

func (cr TransactionCategoryRepository) Browse(search model.BrowseTransactionCategory) ([]model.TransactionCategory, error) {
	var (
		limit        = search.Limit
		offset       = limit * (search.Page - 1)
		categories   []model.TransactionCategory
		sqlStatement = `
			SELECT id, name, description
			FROM transaction_categories
			LIMIT $1
			OFFSET $2
		`
	)

	rows, err := cr.DB.Queryx(sqlStatement, limit, offset)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - Browse : %w", err))
		return categories, err
	}

	for rows.Next() {
		var category model.TransactionCategory
		err := rows.StructScan(&category)
		if err != nil {
			log.Error(fmt.Errorf("error CategoryRepository - Browse : %w", err))
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r TransactionCategoryRepository) UpdateByID(category model.TransactionCategory) error {
	var (
		sqlStatement = `
			UPDATE transaction_categories
			SET updated_at = NOW(),
				name = $2,
				description = $3
			WHERE id = $1
		`
	)

	result, err := r.DB.Exec(sqlStatement, category.ID, category.Name, category.Description)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - UpdateByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}

func (r TransactionCategoryRepository) DeleteByID(id string) error {
	var (
		sqlStatement = `
			DELETE FROM transaction_categories
			WHERE id = $1
		`
	)

	result, err := r.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - DeleteByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}
