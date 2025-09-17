package bot

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"os"
	"strings"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vas-sh/todo/internal/bot/mocks"
	"github.com/vas-sh/todo/internal/models"
	"go.uber.org/mock/gomock"
)

func TestGetPageSuccess(t *testing.T) {
	// arrange
	testCases := []struct {
		name  string
		want  int
		input string
	}{
		{
			name:  "invalid input",
			want:  0,
			input: string(models.NextButtonType) + "3",
		},
		{
			name:  "invalid page type",
			want:  0,
			input: string(models.NextButtonType) + ":x",
		},
		{
			name:  "success",
			want:  5,
			input: string(models.NextButtonType) + ":5",
		},
	}
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})

			//act
			s := New(nil, nil, nil, slog.New(h))
			got := s.getPage(context.Background(), ts.input)

			// assert
			if ts.want != got {
				t.Errorf("want: %d, got: %d", ts.want, got)
			}
		})
	}
}

func TestUpdates(t *testing.T) {
	// arrange
	ctx := context.Background()
	updateCh := make(chan tgbotapi.Update, 1)
	updateCh <- tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: "/starttoken",
			From: &tgbotapi.User{ID: 1},
		},
	}
	close(updateCh)
	ctrl := gomock.NewController(t)
	mockBotSrv := mocks.NewMockboter(ctrl)
	mockBotSrv.EXPECT().GetUpdatesChan(gomock.Any()).Return(updateCh)
	mockBotSrv.EXPECT().Request(gomock.Any()).Return(nil, nil)
	ch := make(chan struct{})
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	s := New(nil, nil, mockBotSrv, slog.New(h))
	mockBotSrv.EXPECT().Send(gomock.Any()).
		DoAndReturn(func(c tgbotapi.Chattable) (tgbotapi.Message, error) {
			sticker := c.(tgbotapi.StickerConfig)
			if string(sticker.File.(tgbotapi.FileID)) != s.sticker.Confused {
				t.Errorf("expected confused sticker, got: %s", sticker.File)
			}
			ch <- struct{}{}
			return tgbotapi.Message{}, nil
		})

	// act
	s.Updates(ctx)
	<-ch
}

func TestCreateMenuFailed(t *testing.T) {
	// arrange
	ctx := context.Background()
	wantErr := errors.New("some error")
	ctrl := gomock.NewController(t)
	mockBotSrv := mocks.NewMockboter(ctrl)
	mockBotSrv.EXPECT().Request(gomock.Any()).Return(nil, wantErr)
	var buf bytes.Buffer
	h := slog.NewTextHandler(&buf, &slog.HandlerOptions{})

	// act
	s := New(nil, nil, mockBotSrv, slog.New(h))
	s.createMenu(ctx)
	err := buf.String()

	// assert
	if !strings.Contains(err, wantErr.Error()) {
		t.Errorf("want: %s, got: %s", wantErr, err)
	}
}

func TestManageUpdateMessageStartCommand(t *testing.T) {
	// arrange
	update := tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: "/starttoken",
			From: &tgbotapi.User{ID: 1},
		},
	}
	wantErr := errors.New("failed to send")
	ctrl := gomock.NewController(t)
	mockBotSrv := mocks.NewMockboter(ctrl)
	mockBotSrv.EXPECT().Send(gomock.Any()).Return(tgbotapi.Message{}, wantErr)
	var buf bytes.Buffer
	h := slog.NewTextHandler(&buf, &slog.HandlerOptions{})

	// act
	s := New(nil, nil, mockBotSrv, slog.New(h))
	s.manageUpdate(update)
	err := buf.String()

	// assert
	if !strings.Contains(err, wantErr.Error()) {
		t.Errorf("want: %s, got: %s", wantErr, err)
	}
}

func TestManageUpdateTextMassageListCommand(t *testing.T) {
	// arrange
	update := tgbotapi.Update{
		UpdateID: 1,
		Message: &tgbotapi.Message{
			Text: string(models.TaskListCommand),
			From: &tgbotapi.User{ID: 1},
		},
	}
	ctrl := gomock.NewController(t)
	wantErr := errors.New("some error")
	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockUserSrv.EXPECT().GetUserID(gomock.Any(), gomock.Any()).Return(int64(0), wantErr).Times(2)
	var buf bytes.Buffer
	h := slog.NewTextHandler(&buf, &slog.HandlerOptions{})

	// act
	s := New(mockUserSrv, nil, nil, slog.New(h))
	s.manageUpdate(update)
	err := buf.String()

	// assert
	if !strings.Contains(err, wantErr.Error()) {
		t.Errorf("want: %s, got: %s", wantErr, err)
	}
}

func TestManageUpdateCallback(t *testing.T) {
	// arrange
	update := tgbotapi.Update{
		CallbackQuery: &tgbotapi.CallbackQuery{
			Data: string(models.NextButtonType) + ":1",
			Message: &tgbotapi.Message{
				Chat:      &tgbotapi.Chat{ID: 1},
				MessageID: 1,
			},
		},
	}
	wantMsg := "You don't have the next task"
	ctrl := gomock.NewController(t)
	mockBotSrv := mocks.NewMockboter(ctrl)
	mockUserSrv := mocks.NewMockuserServicer(ctrl)

	mockUserSrv.EXPECT().GetUserID(gomock.Any(), gomock.Any()).Return(int64(0), models.ErrNotFound)
	var gotMsg string
	mockBotSrv.EXPECT().Send(gomock.Any()).
		Do(func(msg tgbotapi.Chattable) {
			gotMsg = msg.(tgbotapi.EditMessageTextConfig).Text
		}).Return(tgbotapi.Message{}, errors.New("failed to send"))
	mockBotSrv.EXPECT().Request(gomock.Any()).Return(nil, nil)
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})

	// act
	s := New(mockUserSrv, nil, mockBotSrv, slog.New(h))
	s.manageUpdate(update)

	// assert
	if wantMsg != gotMsg {
		t.Errorf("want: %s, got: %s", wantMsg, gotMsg)
	}
}

func TestHandleTaskCallbackEmptyData(t *testing.T) {
	// arrange
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})

	// act
	s := New(nil, nil, nil, slog.New(h))
	callback := &tgbotapi.CallbackQuery{
		Data: "",
		Message: &tgbotapi.Message{
			MessageID: 0,
			Chat: &tgbotapi.Chat{
				ID: 0,
			},
		},
	}
	err := s.handleTaskCallback(context.Background(), callback)

	// assert
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
