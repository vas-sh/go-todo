package integrationtests

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	urlCreate := "http://localhost:8080/create-task"
	body := url.Values{}
	body.Add("title", "Homework")
	body.Add("description", "Write assey")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlCreate, strings.NewReader(body.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("want: %d, got: %d - %s", http.StatusOK, resp.StatusCode, res)
	}
}
