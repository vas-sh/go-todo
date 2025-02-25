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
		name        string
		title       string
		description string
		err         error
	}{
		{
			name:        "ok",
			title:       "Homework",
			description: "need to finish math",
		},
		{
			name:        "empty title",
			description: "need to finish math",
			err:         models.ErrValueEmpty,
		},
	}

	for _, ts := range testCases {
		repoMock := mocks.NewMockrepoer(gomock.NewController(t))
		if ts.err == nil {
			repoMock.EXPECT().Create(gomock.Any(), ts.title, ts.description).Return(models.Task{}, nil)
		}
		s := New(repoMock)
		_, err := s.Create(context.Background(), ts.title, ts.description)
		if err != ts.err {
			t.Errorf("%s - want: %v, got: %v", ts.name, ts.err, err)
		}
	}
}
