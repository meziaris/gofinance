package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/meziaris/gofinance/internal/app/model"
	log "github.com/sirupsen/logrus"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{DB: db}
}

func (r UserRepository) Create(user model.User) error {
	var (
		sqlStatement = `
			INSERT INTO users (username, hashed_password, full_name)
			VALUES ($1, $2, $3)
		`
	)

	_, err := r.DB.Exec(sqlStatement, user.Username, user.HashedPassword, user.FullName)
	if err != nil {
		log.Error("error UserRepository - Create :", err)
		return err
	}

	return nil
}

func (r UserRepository) GetByUsername(username string) (model.User, error) {
	var (
		user         = model.User{}
		sqlStatement = `
			SELECT id, username, full_name, hashed_password
			FROM users
			WHERE username = $1
			LIMIT 1
		`
	)

	if err := r.DB.QueryRowx(sqlStatement, username).StructScan(&user); err != nil {
		log.Error("error UserRepository - GetByUsername :", err)
		return user, err
	}

	return user, nil
}

func (r UserRepository) GetByID(userID int) (model.User, error) {
	var (
		user         = model.User{}
		sqlStatement = `
			SELECT id, username, full_name, created_at
			FROM users
			WHERE id = $1
			LIMIT 1
		`
	)

	if err := r.DB.QueryRowx(sqlStatement, userID).StructScan(&user); err != nil {
		log.Error("error UserRepository - GetByID :", err)
		return user, err
	}

	return user, nil
}
