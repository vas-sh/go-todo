package integrationtests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/vas-sh/todo/internal/models"
)

const taskPath = "tasks"

func parseTaskID(t *testing.T, resp []byte) string {
	t.Helper()
	var task models.Task
	err := json.Unmarshal(resp, &task)
	if err != nil {
		t.Errorf("cannot unmarshal task: %s - %s", resp, err)
	}
	return fmt.Sprint(task.ID)
}

func TestCreateAndDeleteTask(t *testing.T) {
	ctx := context.Background()
	token := signUpAndLogin(t)
	defer userTearDown(t, token)
	body := map[string]string{
		"title":       "Homework",
		"description": "Write assey",
	}
	out, err := json.Marshal(body)
	if err != nil {
		t.Error(err)
		return
	}
	param := requestParam{
		endpoint: taskPath,
		token:    token,
		body:     bytes.NewReader(out),
		method:   http.MethodPost,
	}
	resp := sendRequest(t, ctx, param, http.StatusOK)

	id := parseTaskID(t, resp)
	param = requestParam{
		endpoint: taskPath + "?id=" + id,
		token:    token,
		method:   http.MethodDelete,
	}
	sendRequest(t, ctx, param, http.StatusNoContent)
}
