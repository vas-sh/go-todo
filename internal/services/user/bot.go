package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/vas-sh/todo/internal/models"
	"gorm.io/gorm"
)

func (s *srv) CreateBotActivationToken(ctx context.Context, userID int64) (string, error) {
	token, err := s.repo.FindToken(ctx, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}
	if token != "" {
		return token, nil
	}
	token = uuid.New().String()
	body := models.BotUser{
		UserID: userID,
		Token:  token,
	}
	return token, s.repo.CreateActivationToken(ctx, body)
}

func (s *srv) FindBotUser(ctx context.Context, token string) (models.BotUser, error) {
	return s.repo.FindBotUser(ctx, token)
}

func (s *srv) AddTelegramID(ctx context.Context, userID, telegramID int64) error {
	return s.repo.AddTelegramID(ctx, userID, telegramID)
}

func (s *srv) GetUserID(ctx context.Context, telegramID int64) (int64, error) {
	return s.repo.GetUserID(ctx, telegramID)
}
