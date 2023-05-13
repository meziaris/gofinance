package model

type Category struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type BrowseCategory struct {
	Page  int
	Limit int
}
