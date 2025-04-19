package models

import (
	"time"
)

type TaskStatus struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	TaskID    int64     `json:"-"`
	Task      Task      `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	Status    Status    `json:"status" gorm:"type:task_status"`
}
