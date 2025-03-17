package user

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *srv) SignUp(ctx context.Context, body models.CreateUserBody) (*models.User, error) {
	user, err := s.prepareUser(body)
	if err != nil {
		return nil, err
	}
	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	activationID, err := s.repo.CreateActivation(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	err = s.mail.Send(
		body.Email,
		"Please confirm your email",
		"Click here http://localhost:8100/confirm/"+activationID.String(),
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (*srv) prepareUser(body models.CreateUserBody) (*models.User, error) {
	if body.Name == "" {
		return nil, models.ErrNameEmpty
	}
	if body.Email == "" {
		return nil, models.ErrEmailRequired
	}
	if body.Password == "" {
		return nil, models.ErrPasswordEmpty
	}
	user := models.User{
		Name:  body.Name,
		Email: body.Email,
	}
	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return nil, err
	}
	user.Password = string(password)
	return &user, nil
}
