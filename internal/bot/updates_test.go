package bot

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/vas-sh/todo/internal/models"
)

func TestGetPageSuccess(t *testing.T) {
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
			s := New(nil, nil, nil, slog.New(h))
			got := s.getPage(context.Background(), ts.input)
			if ts.want != got {
				t.Fatalf("wnat: %d, got: %d", ts.want, got)
			}
		})
	}
}
