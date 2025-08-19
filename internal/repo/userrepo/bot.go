package userrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) CreateActivationToken(ctx context.Context, body models.BotUser) error {
	return r.db.WithContext(ctx).Create(body).Error
}

func (r *repo) FindToken(ctx context.Context, userID int64) (string, error) {
	var res models.BotUser
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&res).Error
	return res.Token, err
}

func (r *repo) FindBotUser(ctx context.Context, token string) (res models.BotUser, err error) {
	err = r.db.WithContext(ctx).Where("token = ?", token).First(&res).Error
	return
}

func (r *repo) AddTelegramID(ctx context.Context, userID, telegramID int64) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).Model(models.BotUser{}).Updates(map[string]any{
		"telegram_id": telegramID,
	}).Error
}

func (r *repo) GetUserID(ctx context.Context, telegramID int64) (int64, error) {
	var res models.BotUser
	err := r.db.WithContext(ctx).Where("telegram_id = ?", telegramID).First(&res).Error
	return res.UserID, err
}
