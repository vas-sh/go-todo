package user

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *srv) Login(ctx context.Context, email, password string) (models.User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return models.User{}, models.ErrPasswordIncorrect
	}
	return user, nil
}
