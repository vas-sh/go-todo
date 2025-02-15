package integrationtests

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestCreateAndDelete(t *testing.T) {
	ctx := context.Background()
	body := url.Values{}
	body.Add("title", "Homework")
	body.Add("description", "Write assey")
	param := requestParam{endpoint: "create-task", body: strings.NewReader(body.Encode())}
	resp := sendRequest(t, ctx, param, http.StatusOK)

	id := parseID(t, resp)
	body = url.Values{"id": {id}}
	param = requestParam{endpoint: "delete-task", body: strings.NewReader(body.Encode())}
	sendRequest(t, ctx, param, http.StatusOK)
}
