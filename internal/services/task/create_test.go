package task

import (
	"context"
	"testing"

	"github.com/vas-sh/todo/internal/models"
	"github.com/vas-sh/todo/internal/services/task/mocks"
	"go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		name  string
		input models.Task
		err   error
	}{
		{
			name: "ok",
			input: models.Task{
				Title:       "Homework",
				Description: "need to finish math",
				Status:      models.NewStatus,
			},
		},
		{
			name: "empty title",
			input: models.Task{
				Description: "need to finish math",
				Status:      models.NewStatus,
			},
			err: models.ErrValueEmpty,
		},
	}

	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			repoMock := mocks.NewMockrepoer(gomock.NewController(t))
			if ts.err == nil {
				repoMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			}
			s := New(repoMock)
			_, err := s.Create(context.Background(), ts.input)
			if err != ts.err {
				t.Errorf("want: %v, got: %v", ts.err, err)
			}
		})
	}
}
