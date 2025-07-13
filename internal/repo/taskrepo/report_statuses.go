package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) ReportStatuses(ctx context.Context, userID int64) (models.CountStatus, error) {
	var statuses models.CountStatus
	q := `
		WITH statuses AS (
			SELECT t.status
			FROM tasks AS t
			WHERE t.user_id = ?
		)

		SELECT SUM(CASE WHEN s.status = 'new' THEN 1 END) AS new_status,
			   SUM(CASE WHEN s.status = 'inProgress' THEN 1 END) AS in_progress,
			   SUM(CASE WHEN s.status = 'done' THEN 1 END) AS done,
			   SUM(CASE WHEN s.status = 'canceled' THEN 1 END) AS canceled
		FROM statuses AS s
	`
	err := r.db.WithContext(ctx).Raw(q, userID).First(&statuses).Error
	return statuses, err
}
