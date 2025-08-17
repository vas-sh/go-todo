package bot

import (
	"context"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vas-sh/todo/internal/models"
)

//go:generate mockgen -source=core.go -destination=mocks/mocks.go -package mocks

type userServicer interface {
	FindBotUser(ctx context.Context, token string) (models.BotUser, error)
	AddTelegramID(ctx context.Context, userID, chatID int64) error
	GetUserID(ctx context.Context, chatID int64) (int64, error)
}

type taskServecer interface {
	GetTask(ctx context.Context, userID, index int64) (models.Task, bool, error)
}

type boter interface {
	Request(c tgbotapi.Chattable) (*tgbotapi.APIResponse, error)
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
	GetUpdatesChan(config tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel
}

type stickerFileID struct {
	SuccessStickerFileID string
	FailedStickerFileID  string
}

type srv struct {
	bot     boter
	userSrv userServicer
	taskSrv taskServecer
	logger  *slog.Logger
	sticker stickerFileID
}

func New(userSrv userServicer, taskSrv taskServecer, bot boter, logger *slog.Logger) *srv {
	return &srv{
		bot:     bot,
		userSrv: userSrv,
		taskSrv: taskSrv,
		logger:  logger,
		sticker: stickerFileID{
			SuccessStickerFileID: "CAACAgEAAxkBAAIEbWigGoo5W2lk9Yy7e2hqF3g6EJaTAAKfAwACid9YRM6KQLzK3HtFNgQ",
			FailedStickerFileID:  "CAACAgEAAxkBAAIEeGigIiNK1DJbtmr7j4Y7aXCvpUy5AAL-AgACgSIgRAmiojYO88U7NgQ",
		},
	}
}
