package model

type Currency struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type BrowseCurrency struct {
	Page  int
	Limit int
}
