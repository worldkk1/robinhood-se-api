package domain

import "time"

type Comment struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	UserID    string    `json:"userId"`
	TaskID    string    `json:"taskId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      User      `json:"user"`
}
