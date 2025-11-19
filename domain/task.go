package domain

import "time"

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	UserID      string     `json:"userId"`
	ArchivedAt  *time.Time `json:"archivedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	User        User       `json:"user"`
	TaskLog     []TaskLog  `json:"taskLogs"`
}
