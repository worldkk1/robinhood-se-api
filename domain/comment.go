package domain

import "time"

type Comment struct {
	ID string `json:"id"`
	Content string `json:"content"`
	UserID string `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
