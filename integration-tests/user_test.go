package integrationtests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/vas-sh/todo/internal/config"
	"github.com/vas-sh/todo/internal/db"
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
	token := signUpAndLogin(t)
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

func signUpAndLogin(t *testing.T) string {
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
	activateUser(ctx, t, user.ID)
	token := login(t, models.LoginBody{
		Username: body.Email,
		Password: body.Password,
	})
	return token
}

func getActivationID(t *testing.T, userID int64) (uuid.UUID, error) {
	t.Helper()
	databace, err := db.New(config.Config.DB)
	if err != nil {
		return uuid.Nil, err
	}
	sqlDB, err := databace.DB()
	if err != nil {
		return uuid.Nil, err
	}
	defer sqlDB.Close()
	var activation models.UserActivation
	err = databace.Where("user_id = ? AND date >= now() - interval '1 hour' AND activated = false", userID).
		Order("date DESC").
		First(&activation).Error
	if err != nil {
		return uuid.Nil, err
	}
	return activation.ID, nil
}

func activateUser(ctx context.Context, t *testing.T, userID int64) {
	t.Helper()
	id, err := getActivationID(t, userID)
	if err != nil {
		t.Error(err.Error())
		return
	}
	param := requestParam{
		endpoint: userPath + "/confirm/" + id.String(),
		method:   http.MethodGet,
	}
	sendRequest(t, ctx, param, http.StatusOK)
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
