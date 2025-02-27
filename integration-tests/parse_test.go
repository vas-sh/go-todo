package integrationtests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/vas-sh/todo/internal/models"
)

func parseID(t *testing.T, resp []byte) string {
	t.Helper()
	var task models.Task
	err := json.Unmarshal(resp, &task)
	if err != nil {
		t.Errorf("cannot unmarshal task: %s - %s", resp, err)
	}
	return fmt.Sprint(task.ID)
}
