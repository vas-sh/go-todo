package integrationtests

import (
	"context"
	"io"
	"net/http"
	"testing"
)

type requestParam struct {
	endpoint string
	body     io.Reader
}

func sendRequest(t *testing.T, ctx context.Context, param requestParam, wantStatus int) []byte {
	t.Helper()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:8080/"+param.endpoint, param.body)
	if err != nil {
		t.Error(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
