package models

type User struct {
	ID       int64   `json:"id" gorm:"primary_key"`
	Name     string  `json:"name"`
	Email    *string `json:"email" gorm:"unique"`
	Phone    *int64  `json:"phone" gorm:"unique"`
	Password string  `json:"-"`
}

type CreateUserBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    int64  `json:"phone"`
	Password string `json:"password"`
}
