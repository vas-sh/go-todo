package taskrepo

import "database/sql"

type repo struct {
	db *sql.DB
}

func New(db *sql.DB) *repo {
	return &repo{db: db}
}
