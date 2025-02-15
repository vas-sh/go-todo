package integrationtests

import (
	"strings"
	"testing"
)

func parseID(t *testing.T, resp []byte) string {
	t.Helper()
	items := strings.Split(string(resp), "Homework - Write assey")
	if len(items) != 2 {
		t.Errorf("want: 1 'Homework', got: %d", len(items)-1)
	}
	form := items[len(items)-1]
	ids := strings.Split(form, `value="`)
	if len(ids) != 2 {
		t.Errorf("want: 1 'id', got: %d", len(ids)-1)
		return ""
	}
	id := strings.Split(ids[len(ids)-1], `"`)[0]
	return id
}
