package models

import "errors"

var (
	ErrValueEmpty    = errors.New("value can't be empty")
	ErrNameEmpty     = errors.New("'name' can't be empty")
	ErrEmailRequired = errors.New("'email' is required")
	ErrPasswordEmpty = errors.New("'password' can't be empty")
)
