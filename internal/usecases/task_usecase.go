package usecases

import (
	"time"

	"github.com/worldkk1/robinhood-se-api/domain"
)

type TaskUserDetail struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TaskDetail struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      domain.TaskStatus `json:"status"`
	UserID      string            `json:"userId"`
	ArchivedAt  *time.Time        `json:"archivedAt"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
	User        TaskUserDetail    `json:"user"`
	TaskLogs    []domain.TaskLog  `json:"taskLogs"`
}

type TaskList struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      domain.TaskStatus `json:"status"`
	UserID      string            `json:"userId"`
	ArchivedAt  *time.Time        `json:"archivedAt"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}

type TaskUsecase interface {
	CreateTask(input domain.Task) error
	GetTaskList() []TaskList
	GetTaskDetail(id string) *TaskDetail
	EditTask(id string, input domain.Task, updatedBy string) error
	ArchiveTask(id string, updatedBy string) error
}
