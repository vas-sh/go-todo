package models

type Status string

const (
	NewStatus        Status = "new"
	InProgressStatus Status = "inProgress"
	DoneStatus       Status = "done"
	CanceledStatus   Status = "canceled"
)

type Task struct {
	ID          int64 `gorm:"primary_key"`
	Title       string
	Description string
	Status      Status `gorm:"type:task_status"`
}
