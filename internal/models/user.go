package models

import (
	"time"

	"github.com/google/uuid"
)

const UserContextKey = "user"

type User struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-"`
	Activated bool   `json:"activated"`
}

type CreateUserBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserActivation struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID    int64
	User      User `gorm:"constraint:OnDelete:CASCADE;"`
	Date      time.Time
	Activated bool
}
