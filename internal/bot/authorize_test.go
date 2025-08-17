package bot

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vas-sh/todo/internal/bot/mocks"
	"github.com/vas-sh/todo/internal/models"
	"go.uber.org/mock/gomock"
)

func TestAuthorizeInvalidAuth(t *testing.T) {
	auth := "starttoken123"
	ctrl := gomock.NewController(t)
	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)
	mockTaskSrv := mocks.NewMocktaskServecer(ctrl)

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	s := New(mockUserSrv, mockTaskSrv, mockBotSrv, slog.New(h))
	s.authorize(context.Background(), 0, auth)
}

func TestAuthorizeTokenNotFound(t *testing.T) {
	auth := "start token123"
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)

	mockUserSrv.EXPECT().FindBotUser(context.Background(), "token123").Return(models.BotUser{}, models.ErrNotFound)
	mockBotSrv.EXPECT().Send(gomock.Any()).Return(tgbotapi.Message{}, nil)

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	s := New(mockUserSrv, nil, mockBotSrv, slog.New(h))
	s.authorize(context.Background(), 0, auth)
}

func TestAuthorizeDatabaceError(t *testing.T) {
	auth := "start token123"
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)

	mockUserSrv.EXPECT().FindBotUser(context.Background(), "token123").Return(models.BotUser{}, nil)
	mockUserSrv.EXPECT().AddTelegramID(context.Background(), gomock.Any(), gomock.Any()).Return(errors.New("some error"))
	mockBotSrv.EXPECT().Send(gomock.Any()).Return(tgbotapi.Message{}, nil)

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	s := New(mockUserSrv, nil, mockBotSrv, slog.New(h))
	s.authorize(context.Background(), 0, auth)
}

func TestAuthorizeSuccess(t *testing.T) {
	auth := "start token123"
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)

	mockUserSrv.EXPECT().FindBotUser(context.Background(), "token123").Return(models.BotUser{}, nil)
	mockUserSrv.EXPECT().AddTelegramID(context.Background(), gomock.Any(), gomock.Any()).Return(nil)
	mockBotSrv.EXPECT().Send(gomock.Any()).Return(tgbotapi.Message{}, nil)

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	s := New(mockUserSrv, nil, mockBotSrv, slog.New(h))
	s.authorize(context.Background(), 0, auth)
}
