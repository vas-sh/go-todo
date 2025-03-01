package models

import "errors"

var (
	ErrValueEmpty           = errors.New("value can't be empty")
	ErrNameEmpty            = errors.New("'name' can't be empty")
	ErrEmailOrPhoneRequired = errors.New("'email' or 'phone' is required")
	ErrPasswordEmpty        = errors.New("'password' can't be empty")
)
