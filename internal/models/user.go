package models

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
