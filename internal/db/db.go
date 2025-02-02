package db

import (
	"database/sql"
)

func New(dns string) (*sql.DB, error) {
	return sql.Open("postgres", dns)
}
