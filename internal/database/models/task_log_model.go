package models

import "time"

type TaskLogModel struct {
	ID          string          `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	TaskID      string          `gorm:"column:task_id"`
	Title       *string         `gorm:"column:title"`
	Description *string         `gorm:"column:description"`
	Status      *TaskStatusEnum `gorm:"column:status;type:task_status"`
	CreatedAt   time.Time       `gorm:"column:created_at"`
	UpdatedAt   time.Time       `gorm:"column:updated_at"`
}

func (TaskLogModel) TableName() string {
	return "task_logs"
}
