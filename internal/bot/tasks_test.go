package bot

import (
	"context"
	"log/slog"
	"os"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vas-sh/todo/internal/bot/mocks"
	"github.com/vas-sh/todo/internal/models"
	"go.uber.org/mock/gomock"
)

func TestSendTaskMessageFailedGetTask(t *testing.T) {
	// arrange
	ctx := context.Background()
	wantMsg := "You don't have any tasks yet"
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)
	mockTaskSrv := mocks.NewMocktaskServecer(ctrl)

	mockUserSrv.EXPECT().GetUserID(ctx, gomock.Any()).Return(int64(0), nil)
	mockTaskSrv.EXPECT().GetTask(ctx, gomock.Any(), gomock.Any()).Return(models.Task{}, false, models.ErrNotFound)
	var gotMsg string
	mockBotSrv.EXPECT().Send(gomock.Any()).
		Do(func(msg tgbotapi.Chattable) {
			gotMsg = msg.(tgbotapi.MessageConfig).Text
		}).Return(tgbotapi.Message{}, nil)
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})

	// act
	s := New(mockUserSrv, mockTaskSrv, mockBotSrv, slog.New(h))
	err := s.sendTaskMessage(ctx, 0)

	// assert
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if wantMsg != gotMsg {
		t.Errorf("want: %s, got: %s", wantMsg, gotMsg)
	}
}

func TestSendTaskMessageSuccess(t *testing.T) {
	// arrange
	ctx := context.Background()
	want := "<b>📝 </b>\n"
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)
	mockTaskSrv := mocks.NewMocktaskServecer(ctrl)

	mockUserSrv.EXPECT().GetUserID(ctx, gomock.Any()).Return(int64(0), nil)
	mockTaskSrv.EXPECT().GetTask(ctx, gomock.Any(), gomock.Any()).Return(models.Task{}, true, nil)
	var got string
	mockBotSrv.EXPECT().Send(gomock.Any()).
		Do(func(msg tgbotapi.Chattable) {
			got = msg.(tgbotapi.MessageConfig).Text
		}).Return(tgbotapi.Message{}, nil)

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})

	// act
	s := New(mockUserSrv, mockTaskSrv, mockBotSrv, slog.New(h))
	err := s.sendTaskMessage(ctx, 0)

	// assert
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestRefreshTaskMessageFailedGetTask(t *testing.T) {
	// arrange
	ctx := context.Background()
	want := "You don't have the next task"
	var expectError error
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)
	mockTaskSrv := mocks.NewMocktaskServecer(ctrl)

	mockUserSrv.EXPECT().GetUserID(ctx, gomock.Any()).Return(int64(0), nil)
	mockTaskSrv.EXPECT().GetTask(ctx, gomock.Any(), gomock.Any()).Return(models.Task{}, false, models.ErrNotFound)
	var got string
	mockBotSrv.EXPECT().Send(gomock.Any()).
		Do(func(msg tgbotapi.Chattable) {
			got = msg.(tgbotapi.EditMessageTextConfig).Text
		}).Return(tgbotapi.Message{}, nil)
	h := slog.NewJSONHandler(os.Stdout, nil)

	// act
	s := New(mockUserSrv, mockTaskSrv, mockBotSrv, slog.New(h))
	err := s.refreshTaskMessage(ctx, 0, 0, 0)

	// assert
	if expectError != err {
		t.Errorf("want: %s, got: %s", expectError, err)
	}
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestRefreshTaskMessageSuccess(t *testing.T) {
	// arrange
	ctx := context.Background()
	want := "<b>📝 </b>\n"
	var wantErr error
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)
	mockTaskSrv := mocks.NewMocktaskServecer(ctrl)

	mockUserSrv.EXPECT().GetUserID(ctx, gomock.Any()).Return(int64(0), nil)
	mockTaskSrv.EXPECT().GetTask(ctx, gomock.Any(), gomock.Any()).Return(models.Task{}, true, nil)
	var got string
	mockBotSrv.EXPECT().Send(gomock.Any()).
		Do(func(msg tgbotapi.Chattable) {
			got = msg.(tgbotapi.EditMessageTextConfig).Text
		}).Return(tgbotapi.Message{}, nil)
	h := slog.NewJSONHandler(os.Stdout, nil)

	// act
	s := New(mockUserSrv, mockTaskSrv, mockBotSrv, slog.New(h))
	err := s.refreshTaskMessage(ctx, 0, 0, 0)

	// assert
	if wantErr != err {
		t.Errorf("want: %s, got: %s", wantErr, err)
	}
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestGetTaskFailedGetUserID(t *testing.T) {
	// arrange
	ctx := context.Background()
	wantTask := models.Task{}
	wantNextTaskExist := false
	wantErr := models.ErrNotFound
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockUserSrv.EXPECT().GetUserID(ctx, gomock.Any()).Return(int64(0), models.ErrNotFound)
	h := slog.NewJSONHandler(os.Stdout, nil)

	// act
	s := New(mockUserSrv, nil, nil, slog.New(h))
	gotTask, gotNextTaskExist, gotErr := s.getTask(ctx, 0, 0)

	// assert
	if wantErr.Error() != gotErr.Error() {
		t.Errorf("want: %s, got: %s", wantErr, gotErr)
	}
	if wantNextTaskExist != gotNextTaskExist {
		t.Errorf("want: %v, got: %v", wantNextTaskExist, gotNextTaskExist)
	}
	if wantTask != gotTask {
		t.Errorf("want: %#v, got: %#v", wantTask, gotTask)
	}
}

func TestGetTaskSuccess(t *testing.T) {
	// arrange
	ctx := context.Background()
	wantTask := models.Task{}
	wantNextTaskExist := false
	var wantErr error
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockTaskSrv := mocks.NewMocktaskServecer(ctrl)

	mockUserSrv.EXPECT().GetUserID(ctx, gomock.Any()).Return(int64(0), nil)
	mockTaskSrv.EXPECT().GetTask(ctx, gomock.Any(), gomock.Any()).Return(models.Task{}, false, nil)
	h := slog.NewJSONHandler(os.Stdout, nil)

	// act
	s := New(mockUserSrv, mockTaskSrv, nil, slog.New(h))
	gotTask, gotNextTaskExist, gotErr := s.getTask(ctx, 0, 0)

	// assert
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
