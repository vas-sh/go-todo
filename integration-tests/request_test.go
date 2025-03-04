package integrationtests

import (
	"context"
	"io"
	"net/http"
	"testing"
)

const (
	rootURL = "http://localhost:8080/api/"
)

type requestParam struct {
	endpoint string
	body     io.Reader
	method   string
}

func sendRequest(t *testing.T, ctx context.Context, param requestParam, wantStatus int) []byte {
	t.Helper()
	req, err := http.NewRequestWithContext(ctx, param.method, rootURL+param.endpoint, param.body)
	if err != nil {
		t.Error(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		return nil
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return nil
	}
	if resp.StatusCode != wantStatus {
		t.Errorf("want: %d, got: %d - %s", wantStatus, resp.StatusCode, res)
		return nil
	}
	return res
}
