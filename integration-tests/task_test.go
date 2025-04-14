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
		t.Error(err.Error())
	}
	return fmt.Sprint(task.ID)
}

func deleteTask(ctx context.Context, t *testing.T, token string, resp []byte) {
	t.Helper()
	id := parseTaskID(t, resp)
	param := requestParam{
		endpoint: taskPath + "/" + id,
		token:    token,
		method:   http.MethodDelete,
	}
	sendRequest(t, ctx, param, http.StatusNoContent)
}

func updateTask(ctx context.Context, t *testing.T, token, id string) {
	t.Helper()
	body := map[string]any{
		"title":       "Homework 1",
		"description": "Write assey 1",
		"status":      models.DoneStatus,
	}
	out, err := json.Marshal(body)
	if err != nil {
		t.Error(err)
		return
	}

	param := requestParam{
		endpoint: taskPath + "/" + id,
		token:    token,
		body:     bytes.NewReader(out),
		method:   http.MethodPut,
	}
	sendRequest(t, ctx, param, http.StatusOK)
}

func createTask(ctx context.Context, t *testing.T, token string) []byte {
	t.Helper()
	body := map[string]string{
		"title":       "Homework",
		"description": "Write assey",
	}
	out, err := json.Marshal(body)
	if err != nil {
		t.Error(err)
		return nil
	}
	param := requestParam{
		endpoint: taskPath,
		token:    token,
		body:     bytes.NewReader(out),
		method:   http.MethodPost,
	}
	return sendRequest(t, ctx, param, http.StatusOK)
}

func TestCreateAndDeleteTask(t *testing.T) {
	ctx := context.Background()
	token := signUpAndLogin(t)
	defer userTearDown(t, token)
	resp := createTask(ctx, t, token)
	deleteTask(ctx, t, token, resp)
}

func TestUpdateTask(t *testing.T) {
	ctx := context.Background()
	token := signUpAndLogin(t)
	defer userTearDown(t, token)

	resp := createTask(ctx, t, token)
	id := parseTaskID(t, resp)
	defer deleteTask(ctx, t, token, resp)
	updateTask(ctx, t, token, id)

	param := requestParam{
		endpoint: taskPath,
		token:    token,
		method:   http.MethodGet,
	}
	resp = sendRequest(t, ctx, param, http.StatusOK)
	var tasks []models.Task
	err := json.Unmarshal(resp, &tasks)
	if err != nil {
		t.Error(err.Error())
	}
	for i := range tasks {
		if fmt.Sprint(tasks[i].ID) != id {
			continue
		}
		want := models.Task{
			ID:          tasks[i].ID,
			Title:       "Homework 1",
			Description: "Write assey 1",
			Status:      models.DoneStatus,
		}
		if want != tasks[i] {
			t.Errorf("want: %#v, got: %#v", want, tasks[i])
		}
		return
	}
	t.Errorf("task with id %s is not found", id)
}
