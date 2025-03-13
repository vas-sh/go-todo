package integrationtests

import (
	"bytes"
	"context"
	"encoding/json"
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
	token := createUserAndLogin(t)
	userTearDown(t, token)
}

func login(t *testing.T, body models.LoginBody) string {
	out, err := json.Marshal(body)
	if err != nil {
		t.Error(err)
		return ""
	}

	param := requestParam{
		endpoint: userPath + "/login",
		body:     bytes.NewReader(out),
		method:   http.MethodPost,
	}
	resp := sendRequest(t, context.Background(), param, http.StatusOK)
	var jwtToken struct {
		Token string `json:"token"`
		Type  string `json:"type"`
	}
	err = json.Unmarshal(resp, &jwtToken)
	if err != nil {
		t.Error(err)
		return ""
	}
	if jwtToken.Type != "JWT" {
		t.Errorf("want: 'JWT', got: %q", jwtToken.Type)
	}
	if jwtToken.Token == "" {
		t.Errorf("got empty token")
	}
	return jwtToken.Type + " " + jwtToken.Token
}

func createUserAndLogin(t *testing.T) string {
	t.Helper()
	ctx := context.Background()
	body := models.CreateUserBody{
		Name:     "Jhon",
		Email:    "john@gmail.com",
		Password: "1111",
	}
	out, err := json.Marshal(body)
	if err != nil {
		t.Error(err)
		return ""
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

	token := login(t, models.LoginBody{
		Username: body.Email,
		Password: body.Password,
	})
	return token
}

func userTearDown(t *testing.T, token string) {
	t.Helper()
	param := requestParam{
		endpoint: userPath,
		method:   http.MethodDelete,
		token:    token,
	}
	sendRequest(t, context.Background(), param, http.StatusNoContent)
}
