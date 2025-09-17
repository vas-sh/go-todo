package bot

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"testing"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vas-sh/todo/internal/bot/mocks"
	"github.com/vas-sh/todo/internal/models"
	"go.uber.org/mock/gomock"
)

func TestSendTaskMessageNotFound(t *testing.T) {
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

func TestSendTaskMessageGotError(t *testing.T) {
	ctx := context.Background()
	wantErr := errors.New("some error")
	ctrl := gomock.NewController(t)

	mockUserSrv := mocks.NewMockuserServicer(ctrl)
	mockBotSrv := mocks.NewMockboter(ctrl)
	mockTaskSrv := mocks.NewMocktaskServecer(ctrl)

	mockUserSrv.EXPECT().GetUserID(ctx, gomock.Any()).Return(int64(0), nil)
	mockTaskSrv.EXPECT().GetTask(ctx, gomock.Any(), gomock.Any()).Return(models.Task{}, false, wantErr)
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})

	// act
	s := New(mockUserSrv, mockTaskSrv, mockBotSrv, slog.New(h))
	err := s.sendTaskMessage(ctx, 0)

	// assert
	if !errors.Is(err, wantErr) {
		t.Errorf("unexpected error: %s", err)
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
		t.Errorf("want: %s, got: %s", wantErr, gotErr)
	}
	if wantNextTaskExist != gotNextTaskExist {
		t.Errorf("want: %v, got: %v", wantNextTaskExist, gotNextTaskExist)
	}
	if wantTask != gotTask {
		t.Errorf("want: %#v, got: %#v", wantTask, gotTask)
	}
}

func TestTaskFormat(t *testing.T) {
	// arrange
	time := time.Now()
	testCases := []struct {
		name string
		task models.Task
		want string
	}{
		{
			name: "task status new",
			task: models.Task{
				Title:        "task 1",
				Description:  "description 1",
				Status:       models.NewStatus,
				EstimateTime: &time,
			},
			want: "<b>📝 task 1</b>\n📎 <i>description 1</i>\n🟡 New\nDeadline: " +
				time.Format("02/01/2006 15:04") +
				" (-0 days)",
		},
		{
			name: "task status done",
			task: models.Task{
				Title:        "task 2",
				Description:  "description 2",
				Status:       models.DoneStatus,
				EstimateTime: &time,
			},
			want: "<b>📝 task 2</b>\n📎 <i>description 2</i>\n🟢 Done\nDeadline: " +
				time.Format("02/01/2006 15:04") +
				" (-0 days)",
		},
		{
			name: "task status in progress",
			task: models.Task{
				Title:        "task 3",
				Description:  "description 3",
				Status:       models.InProgressStatus,
				EstimateTime: &time,
			},
			want: "<b>📝 task 3</b>\n📎 <i>description 3</i>\n🟡 In progress\nDeadline: " +
				time.Format("02/01/2006 15:04") +
				" (-0 days)",
		},
		{
			name: "task status canceled",
			task: models.Task{
				Title:        "task 4",
				Description:  "description 4",
				Status:       models.CanceledStatus,
				EstimateTime: &time,
			},
			want: "<b>📝 task 4</b>\n📎 <i>description 4</i>\n🔴 Canceled\nDeadline: " +
				time.Format("02/01/2006 15:04") +
				" (-0 days)",
		},
	}
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			h := slog.NewJSONHandler(os.Stdout, nil)

			// act
			s := New(nil, nil, nil, slog.New(h))
			got := s.taskFormat(ts.task)

			// assert
			if ts.want != got {
				t.Errorf("want: %s, got: %s", ts.want, got)
			}
		})
	}
}

func TestKeyboard(t *testing.T) {
	// arrange
	testCases := []struct {
		name string
		page int
		want tgbotapi.InlineKeyboardMarkup
		ops  keyboardOps
	}{
		{
			name: "Prev and Next buttons",
			page: 5,
			want: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("⬅️ Prev", fmt.Sprintf("%s:%d", models.PrevButtonType, 5-1)),
					tgbotapi.NewInlineKeyboardButtonData("➡️ Next", fmt.Sprintf("%s:%d", models.NextButtonType, 5+1)),
				),
			),
			ops: keyboardOps{
				exist: true,
			},
		},
		{
			name: "Prev button",
			page: 3,
			want: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("⬅️ Prev", fmt.Sprintf("%s:%d", models.PrevButtonType, 3-1)),
				),
			),
			ops: keyboardOps{
				exist: false,
			},
		},
		{
			name: "Next button",
			page: 0,
			want: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("➡️ Next", fmt.Sprintf("%s:%d", models.NextButtonType, 0+1)),
				),
			),
			ops: keyboardOps{
				exist: true,
			},
		},
		{
			name: "Empty keyboard",
			page: 0,
			want: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{}},
			ops: keyboardOps{
				exist: false,
			},
		},
	}
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			h := slog.NewJSONHandler(os.Stdout, nil)

			// act
			s := New(nil, nil, nil, slog.New(h))
			got := s.listKeyboard(ts.page, ts.ops)

			// assert
			if !reflect.DeepEqual(ts.want, got) {
				t.Errorf("want: %#v, got: %#v", ts.want, got)
			}
		})
	}
}
