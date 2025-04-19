package integrationtests

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/vas-sh/todo/internal/models"
)

func TestStatusChange(t *testing.T) {
	ctx := context.Background()
	token := signUpAndLogin(t)
	defer userTearDown(t, token)

	resp := createTask(ctx, t, token)
	id := parseTaskID(t, resp)
	defer deleteTask(ctx, t, token, resp)
	updateTask(ctx, t, token, id)

	param := requestParam{
		endpoint: taskPath + "/" + id + "/statuses",
		token:    token,
		method:   http.MethodGet,
	}
	resp = sendRequest(t, ctx, param, http.StatusOK)
	var statuses []models.TaskStatus
	err := json.Unmarshal(resp, &statuses)
	if err != nil {
		t.Error(err.Error())
	}
	if len(statuses) != 2 {
		t.Errorf("want 2 statuses, got %+v", statuses)
		return
	}
	if statuses[0].Status != models.NewStatus {
		t.Errorf("want 'new' status, got %q", statuses[0].Status)
	}
	if statuses[1].Status != models.DoneStatus {
		t.Errorf("want 'done' status, got %q", statuses[1].Status)
	}
}
