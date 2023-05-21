package model

import "time"

type User struct {
	ID             int       `db:"id"`
	Username       string    `db:"username"`
	FullName       string    `db:"full_name"`
	HashedPassword string    `db:"hashed_password"`
	CreatedAt      time.Time `db:"created_at"`
}
