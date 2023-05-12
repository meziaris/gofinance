package model

type User struct {
	ID             int    `db:"id"`
	Username       string `db:"username"`
	FullName       string `db:"full_name"`
	HashedPassword string `db:"hashed_password"`
}
