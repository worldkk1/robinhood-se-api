package models

import "time"

type CommentModel struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	Content   string    `gorm:"column:content"`
	UserID    string    `gorm:"column:user_id"`
	TaskID    string    `gorm:"column:task_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	User      UserModel `gorm:"foreignKey:UserID;references:ID"`
}

func (CommentModel) TableName() string {
	return "comments"
}
