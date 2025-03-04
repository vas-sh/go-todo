package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vas-sh/todo/internal/models"
)

func TestPrepareUser(t *testing.T) {
	testCases := []struct {
		name string
		body models.CreateUserBody
		want *models.User
		err  error
	}{
		{
			name: "ok",
			body: models.CreateUserBody{
				Name:     "John",
				Email:    "test.email@gmail.com",
				Password: "somepassword",
			},
			want: &models.User{
				Name:  "John",
				Email: "test.email@gmail.com",
			},
		},
		{
			name: "ok; only email",
			body: models.CreateUserBody{
				Name:     "John",
				Email:    "test.email@gmail.com",
				Password: "somepassword",
			},
			want: &models.User{
				Name:  "John",
				Email: "test.email@gmail.com",
			},
		},
		{
			name: "name is missing",
			body: models.CreateUserBody{
				Email:    "test.email@gmail.com",
				Password: "somepassword",
			},
			err: models.ErrNameEmpty,
		},
		{
			name: "password is missing",
			body: models.CreateUserBody{
				Name:  "John",
				Email: "test.email@gmail.com",
			},
			err: models.ErrPasswordEmpty,
		},
		{
			name: "email is missing",
			body: models.CreateUserBody{
				Name: "John",
			},
			err: models.ErrEmailRequired,
		},
	}
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			s := New(nil)
			got, err := s.prepareUser(ts.body)
			if err != ts.err {
				t.Errorf("want: %v, got: %v", ts.err, err)
			}
			if got != nil {
				if got.Password == "" {
					t.Error("want crypto password")
				}
				if got.Password == ts.body.Password {
					t.Error("password is not hashed")
				}
				got.Password = ""
			}
			if diff := cmp.Diff(got, ts.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
