package models

type Status string

const (
	NewStatus        Status = "new"
	InProgressStatus Status = "inProgress"
	DoneStatus       Status = "done"
	CanceledStatus   Status = "canceled"
)

type Task struct {
	ID          int64  `gorm:"primary_key" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `gorm:"type:task_status" json:"status"`
}
