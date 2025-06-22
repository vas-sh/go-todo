package models

import "time"

type Status string

const (
	NewStatus        Status = "new"
	InProgressStatus Status = "inProgress"
	DoneStatus       Status = "done"
	CanceledStatus   Status = "canceled"
)

type CountCompletion struct {
	InTime        int64 `json:"in_time"`
	NotInTime     int64 `json:"not_in_time"`
	ActiveOverdue int64 `json:"active_overdue"`
	DeadLineSoon  int64 `json:"dead_line_soon"`
}

type CountStatus struct {
	Status Status `json:"status"`
	Count  int64  `json:"count"`
}

type Task struct {
	ID           int64      `json:"id" gorm:"primary_key"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Status       Status     `json:"status" gorm:"type:task_status"`
	UserID       int64      `json:"-"`
	User         User       `gorm:"constraint:OnDelete:CASCADE;"`
	EstimateTime *time.Time `json:"estimateTime"`
}
