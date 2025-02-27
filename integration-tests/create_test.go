package integrationtests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateAndDelete(t *testing.T) {
	ctx := context.Background()
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
		body:     bytes.NewReader(out),
		method:   http.MethodPost,
	}
	resp := sendRequest(t, ctx, param, http.StatusOK)

	id := parseID(t, resp)
	param = requestParam{
		endpoint: taskPath + "?id=" + id,
		method:   http.MethodDelete,
	}
	sendRequest(t, ctx, param, http.StatusNoContent)
}
