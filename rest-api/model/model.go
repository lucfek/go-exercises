package model

import "database/sql"

// Model struct
type Model struct {
	db *sql.DB
}

// New gets address of databas as parameter  od returns new Model struct
func New(db *sql.DB) Model {
	return Model{db: db}
}
