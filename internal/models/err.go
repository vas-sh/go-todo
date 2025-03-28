package models

import "errors"

var (
	ErrValueEmpty        = errors.New("value can't be empty")
	ErrNameEmpty         = errors.New("'name' can't be empty")
	ErrEmailRequired     = errors.New("'email' is required")
	ErrPasswordEmpty     = errors.New("'password' can't be empty")
	ErrPasswordIncorrect = errors.New("'password' is invailid")
	ErrInvalidToken      = errors.New("invalid token")
	ErrInvalidUser       = errors.New("invalid user")
	ErrAlreadyExist      = errors.New("already exist")
	ErrAlreadyActivated  = errors.New("user already activated")
	ErrAlreadyExpired    = errors.New("already expired")
)
