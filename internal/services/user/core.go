package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/vas-sh/todo/internal/models"
)

type mailer interface {
	Send(to, subject, body string) error
}

type repoer interface {
	Create(ctx context.Context, user *models.User) error
	Remove(ctx context.Context, id int64) error
	GetByEmail(ctx context.Context, email string) (user models.User, err error)
	CreateActivation(ctx context.Context, userID int64) (uuid.UUID, error)
	Activation(ctx context.Context, id uuid.UUID) (*models.UserActivation, error)
	Activate(ctx context.Context, userActivation *models.UserActivation) error
	CreateActivationToken(ctx context.Context, body models.BotUser) error
	FindBotUser(ctx context.Context, token string) (res models.BotUser, err error)
	AddTelegramID(ctx context.Context, userID, telegramID int64) error
	FindToken(ctx context.Context, userID int64) (string, error)
	GetUserID(ctx context.Context, telegramID int64) (int64, error)
}

type srv struct {
	repo repoer
	mail mailer
}

func New(repo repoer, mail mailer) *srv {
	return &srv{
		repo: repo,
		mail: mail,
	}
}
