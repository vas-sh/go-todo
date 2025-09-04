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
	// arrange
	auth := "starttoken123"
	ctrl := gomock.NewController(t)
	mockBotSrv := mocks.NewMockboter(ctrl)
	h := slog.NewJSONHandler(os.Stdout, nil)
	s := New(nil, nil, mockBotSrv, slog.New(h))
	mockBotSrv.EXPECT().Send(gomock.Any()).
		DoAndReturn(func(c tgbotapi.Chattable) (tgbotapi.Message, error) {
			sticker := c.(tgbotapi.StickerConfig)
			if string(sticker.File.(tgbotapi.FileID)) != s.sticker.Confused {
				t.Errorf("expected confused sticker, got: %s", sticker.File)
			}
			return tgbotapi.Message{}, nil
		})

	// act
	err := s.authorize(context.Background(), 0, auth)

	// assert
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestAuthorizeTokenNotFound(t *testing.T) {
	// arrange
	ctx := context.Background()
	auth := "start token123"
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)

	mockUserSrv.EXPECT().FindBotUser(ctx, "token123").Return(models.BotUser{}, models.ErrNotFound)
	h := slog.NewJSONHandler(os.Stdout, nil)
	s := New(mockUserSrv, nil, mockBotSrv, slog.New(h))
	mockBotSrv.EXPECT().Send(gomock.Any()).
		DoAndReturn(func(c tgbotapi.Chattable) (tgbotapi.Message, error) {
			sticker := c.(tgbotapi.StickerConfig)
			if string(sticker.File.(tgbotapi.FileID)) != s.sticker.Confused {
				t.Errorf("expected confused sticker, got: %s", sticker.File)
			}
			return tgbotapi.Message{}, nil
		}).
		Return(tgbotapi.Message{}, nil)

	// act
	err := s.authorize(ctx, 0, auth)

	// assert
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestAuthorizeDatabaceError(t *testing.T) {
	// arrange
	ctx := context.Background()
	auth := "start token123"
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)

	mockUserSrv.EXPECT().FindBotUser(ctx, "token123").Return(models.BotUser{}, nil)
	mockUserSrv.EXPECT().AddTelegramID(ctx, gomock.Any(), gomock.Any()).Return(errors.New("databace error"))
	h := slog.NewJSONHandler(os.Stdout, nil)
	s := New(mockUserSrv, nil, mockBotSrv, slog.New(h))
	mockBotSrv.EXPECT().Send(gomock.Any()).
		DoAndReturn(func(c tgbotapi.Chattable) (tgbotapi.Message, error) {
			sticker := c.(tgbotapi.StickerConfig)
			if string(sticker.File.(tgbotapi.FileID)) != s.sticker.Confused {
				t.Errorf("expected confused sticker, got: %s", sticker.File)
			}
			return tgbotapi.Message{}, nil
		}).
		Return(tgbotapi.Message{}, nil)

	// act
	err := s.authorize(ctx, 0, auth)

	// assert
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestAuthorizeSuccess(t *testing.T) {
	// arrange
	ctx := context.Background()
	auth := "start token123"
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)

	mockUserSrv.EXPECT().FindBotUser(ctx, "token123").Return(models.BotUser{}, nil)
	mockUserSrv.EXPECT().AddTelegramID(ctx, gomock.Any(), gomock.Any()).Return(nil)
	h := slog.NewJSONHandler(os.Stdout, nil)
	s := New(mockUserSrv, nil, mockBotSrv, slog.New(h))
	mockBotSrv.EXPECT().Send(gomock.Any()).
		DoAndReturn(func(c tgbotapi.Chattable) (tgbotapi.Message, error) {
			sticker := c.(tgbotapi.StickerConfig)
			if string(sticker.File.(tgbotapi.FileID)) != s.sticker.Congratulations {
				t.Errorf("expected confused sticker, got: %s", sticker.File)
			}
			return tgbotapi.Message{}, nil
		}).
		Return(tgbotapi.Message{}, nil)

	// act
	err := s.authorize(ctx, 0, auth)

	// assert
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
