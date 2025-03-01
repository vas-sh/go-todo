package user

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *srv) Create(ctx context.Context, body models.CreateUserBody) (*models.User, error) {
	user, err := s.prepareUser(body)
	if err != nil {
		return nil, err
	}
	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *srv) prepareUser(body models.CreateUserBody) (*models.User, error) {
	if body.Name == "" {
		return nil, models.ErrNameEmpty
	}
	if body.Email == "" && body.Phone == 0 {
		return nil, models.ErrEmailOrPhoneRequired
	}
	if body.Password == "" {
		return nil, models.ErrPasswordEmpty
	}
	user := models.User{
		Name: body.Name,
	}
	if body.Email != "" {
		user.Email = &body.Email
	}
	if body.Phone != 0 {
		user.Phone = &body.Phone
	}
	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return nil, err
	}
	user.Password = string(password)
	return &user, nil
}
