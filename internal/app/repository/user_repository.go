package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meziaris/gofinance/internal/app/model"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{DB: db}
}

func (r UserRepository) Create(user model.UserRepository) error {
	var (
		sqlStatement = `
			INSERT INTO users (username, hashed_password, full_name)
			VALUES ($1, $2, $3)
		`
	)

	_, err := r.DB.Exec(sqlStatement, user.Username, user.HashedPassword, user.FullName)
	if err != nil {
		fmt.Println("error UserRepository - Create :", err)
		return err
	}

	return nil
}

func (r UserRepository) GetByUsername(username string) (model.UserRepository, error) {
	var (
		user         = model.UserRepository{}
		sqlStatement = `
			SELECT id, username, full_name
			FROM users
			WHERE username = $1
			LIMIT 1
		`
	)

	if err := r.DB.QueryRowx(sqlStatement, username).StructScan(&user); err != nil {
		fmt.Println("error UserRepository - GetByUsername :", err)
		return user, err
	}

	return user, nil
}
