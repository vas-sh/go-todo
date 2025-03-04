package integrationtests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vas-sh/todo/internal/models"
)

const userPath = "users"

func parseUser(t *testing.T, resp []byte) models.User {
	t.Helper()
	var user models.User
	err := json.Unmarshal(resp, &user)
	if err != nil {
		t.Errorf("cannot unmarshal user: %s - %s", resp, err)
	}
	return user
}

func TestSignUp(t *testing.T) {
	ctx := context.Background()
	body := models.CreateUserBody{
		Name:     "Jhon",
		Email:    "john@gmail.com",
		Password: "1111",
	}
	out, err := json.Marshal(body)
	if err != nil {
		t.Error(err)
		return
	}
	param := requestParam{
		endpoint: userPath + "/sign-up",
		body:     bytes.NewReader(out),
		method:   http.MethodPost,
	}
	resp := sendRequest(t, ctx, param, http.StatusOK)
	user := parseUser(t, resp)
	want := models.User{
		Name:  body.Name,
		Email: body.Email,
		ID:    user.ID,
	}
	if user != want {
		t.Error(cmp.Diff(user, want))
	}

	param = requestParam{
		endpoint: userPath + "?id=" + fmt.Sprint((user.ID)),
		method:   http.MethodDelete,
	}
	sendRequest(t, ctx, param, http.StatusNoContent)
}
