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

func TestSendTaskMessageFailedGetTask(t *testing.T) {
	want := "You don't have any tasks yet"
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)
	mockTaskSrv := mocks.NewMocktaskServecer(ctrl)

	mockUserSrv.EXPECT().GetUserID(context.Background(), gomock.Any()).Return(int64(0), nil)
	mockTaskSrv.EXPECT().GetTask(context.Background(), gomock.Any(), gomock.Any()).Return(models.Task{}, false, models.ErrNotFound)
	var got string
	mockBotSrv.EXPECT().Send(gomock.Any()).
		Do(func(msg tgbotapi.Chattable) {
			got = msg.(tgbotapi.MessageConfig).Text
		}).Return(tgbotapi.Message{}, nil)

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	s := New(mockUserSrv, mockTaskSrv, mockBotSrv, slog.New(h))
	s.sendTaskMessage(context.Background(), 0)
	if want != got {
		t.Fatalf("want: %s, got: %s", want, got)
	}
}

func TestSendTaskMessageSuccess(t *testing.T) {
	want := "<b>📝 </b>\n"
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)
	mockTaskSrv := mocks.NewMocktaskServecer(ctrl)

	mockUserSrv.EXPECT().GetUserID(context.Background(), gomock.Any()).Return(int64(0), nil)
	mockTaskSrv.EXPECT().GetTask(context.Background(), gomock.Any(), gomock.Any()).Return(models.Task{}, true, nil)
	var got string
	mockBotSrv.EXPECT().Send(gomock.Any()).
		Do(func(msg tgbotapi.Chattable) {
			got = msg.(tgbotapi.MessageConfig).Text
		}).Return(tgbotapi.Message{}, nil)

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	s := New(mockUserSrv, mockTaskSrv, mockBotSrv, slog.New(h))
	s.sendTaskMessage(context.Background(), 0)
	if want != got {
		t.Fatalf("want: %s, got: %s", want, got)
	}
}

func TestRefreshTaskMessageFailedGetTask(t *testing.T) {
	want := "You don't have the next task"
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)
	mockTaskSrv := mocks.NewMocktaskServecer(ctrl)

	mockUserSrv.EXPECT().GetUserID(context.Background(), gomock.Any()).Return(int64(0), nil)
	mockTaskSrv.EXPECT().GetTask(context.Background(), gomock.Any(), gomock.Any()).Return(models.Task{}, false, errors.New("not found"))
	var got string
	mockBotSrv.EXPECT().Send(gomock.Any()).
		Do(func(msg tgbotapi.Chattable) {
			got = msg.(tgbotapi.EditMessageTextConfig).Text
		}).Return(tgbotapi.Message{}, nil)

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	s := New(mockUserSrv, mockTaskSrv, mockBotSrv, slog.New(h))
	s.refreshTaskMessage(context.Background(), 0, 0, 0)
	if want != got {
		t.Fatalf("want: %s, got: %s", want, got)
	}
}

func TestRefreshTaskMessageSuccess(t *testing.T) {
	want := "<b>📝 </b>\n"
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)
	mockTaskSrv := mocks.NewMocktaskServecer(ctrl)

	mockUserSrv.EXPECT().GetUserID(context.Background(), gomock.Any()).Return(int64(0), nil)
	mockTaskSrv.EXPECT().GetTask(context.Background(), gomock.Any(), gomock.Any()).Return(models.Task{}, true, nil)
	var got string
	mockBotSrv.EXPECT().Send(gomock.Any()).
		Do(func(msg tgbotapi.Chattable) {
			got = msg.(tgbotapi.EditMessageTextConfig).Text
		}).Return(tgbotapi.Message{}, nil)

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	s := New(mockUserSrv, mockTaskSrv, mockBotSrv, slog.New(h))
	s.refreshTaskMessage(context.Background(), 0, 0, 0)
	if want != got {
		t.Fatalf("want: %s, got: %s", want, got)
	}
}

func TestGetTaskFailedGetUserID(t *testing.T) {
	wantTask := models.Task{}
	wantNextTaskExist := false
	wantErr := models.ErrNotFound
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockUserSrv.EXPECT().GetUserID(context.Background(), gomock.Any()).Return(int64(0), models.ErrNotFound)

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	s := New(mockUserSrv, nil, nil, slog.New(h))
	gotTask, gotNextTaskExist, gotErr := s.getTask(context.Background(), 0, 0)
	if wantErr.Error() != gotErr.Error() {
		t.Fatalf("want: %s, got: %s", wantErr, gotErr)
	}
	if wantNextTaskExist != gotNextTaskExist {
		t.Fatalf("want: %v, got: %v", wantNextTaskExist, gotNextTaskExist)
	}
	if wantTask != gotTask {
		t.Fatalf("want: %#v, got: %#v", wantTask, gotTask)
	}
}

func TestGetTaskSuccess(t *testing.T) {
	wantTask := models.Task{}
	wantNextTaskExist := false
	var wantErr error
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockTaskSrv := mocks.NewMocktaskServecer(ctrl)

	mockUserSrv.EXPECT().GetUserID(context.Background(), gomock.Any()).Return(int64(0), nil)
	mockTaskSrv.EXPECT().GetTask(context.Background(), gomock.Any(), gomock.Any()).Return(models.Task{}, false, nil)

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	s := New(mockUserSrv, mockTaskSrv, nil, slog.New(h))
	gotTask, gotNextTaskExist, gotErr := s.getTask(context.Background(), 0, 0)
	if wantErr != gotErr {
		t.Fatalf("want: %s, got: %s", wantErr, gotErr)
	}
	if wantNextTaskExist != gotNextTaskExist {
		t.Fatalf("want: %v, got: %v", wantNextTaskExist, gotNextTaskExist)
	}
	if wantTask != gotTask {
		t.Fatalf("want: %#v, got: %#v", wantTask, gotTask)
	}
}
