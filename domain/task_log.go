package domain

import "time"

type TaskLog struct {
	ID          string      `json:"id"`
	TaskID      string      `json:"taskId"`
	Title       *string     `json:"title"`
	Description *string     `json:"description"`
	Status      *TaskStatus `json:"status"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}
