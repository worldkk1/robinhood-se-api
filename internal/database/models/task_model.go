package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/worldkk1/robinhood-se-api/domain"
)

type TaskStatusEnum domain.TaskStatus

const (
	ToDo       TaskStatusEnum = TaskStatusEnum(domain.ToDo)
	InProgress TaskStatusEnum = TaskStatusEnum(domain.InProgress)
	Done       TaskStatusEnum = TaskStatusEnum(domain.Done)
)

func (p *TaskStatusEnum) Scan(value any) error {
	if value == nil {
		*p = TaskStatusEnum("")
		return nil
	}
	switch v := value.(type) {
	case []byte:
		*p = TaskStatusEnum(string(v))
		return nil
	case string:
		*p = TaskStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("unsupported Scan type for TaskStatusEnum: %T", value)
	}
}

func (p TaskStatusEnum) Value() (driver.Value, error) {
	return string(p), nil
}

type TaskModel struct {
	ID          string         `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	Title       string         `gorm:"column:title"`
	Description string         `gorm:"column:description"`
	Status      TaskStatusEnum `gorm:"column:status;type:task_status;default:to_do"`
	UserID      string         `gorm:"column:user_id"`
	ArchivedAt  *time.Time     `gorm:"column:archived_at"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	User        UserModel      `gorm:"foreignKey:UserID;references:ID"`
	TaskLogs    []TaskLogModel `gorm:"foreignKey:TaskID;references:ID"`
}

func (TaskModel) TableName() string {
	return "tasks"
}
