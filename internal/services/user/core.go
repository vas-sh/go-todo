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
