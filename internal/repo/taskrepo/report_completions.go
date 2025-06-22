package taskrepo

import (
	"context"

	"github.com/vas-sh/todo/internal/models"
)

func (r *repo) ReportCompletions(ctx context.Context, userID int64) (models.CountCompletion, error) {
	var completions models.CountCompletion
	q := `	
		WITH temp_tasks AS (
			SELECT t.id, t.status, t.estimate_time,
			 	   (SELECT MAX(s.created_at) 
					FROM task_statuses AS s
					WHERE s.task_id = t.id) AS status_update
			FROM tasks AS t
			WHERE t.user_id = ?
		)

		SELECT SUM(CASE WHEN t.status = 'done' AND t.status_update <= t.estimate_time THEN 1 END) AS in_time,
		 	   SUM(CASE WHEN t.status = 'done' AND t.status_update > t.estimate_time THEN 1 END) AS not_in_time,
			   SUM(CASE WHEN t.status IN ('new', 'inProgress') AND t.estimate_time < NOW() THEN 1 
			   	   END) AS active_overdue,
			   SUM(CASE WHEN t.status IN ('new', 'inProgress') 
			   				 AND t.estimate_time BETWEEN NOW() 
							 AND NOW() + INTERVAL '24 HOURS' THEN 1 
				   END) AS dead_line_soon
		FROM temp_tasks AS t
	`
	err := r.db.WithContext(ctx).Raw(q, userID).First(&completions).Error
	return completions, err
}
