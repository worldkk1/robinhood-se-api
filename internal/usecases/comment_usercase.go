package usecases

import (
	"time"

	"github.com/worldkk1/robinhood-se-api/domain"
)

type CommentUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TaskComment struct {
	ID        string      `json:"id"`
	Content   string      `json:"content"`
	UserID    string      `json:"userId"`
	TaskID    string      `json:"taskId"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	User      CommentUser `json:"user"`
}

type CommentUsecase interface {
	CreateComment(input domain.Comment) error
	GetTaskComments(taskId string) []TaskComment
	EditComment(id string, content string, userId string) error
	DeleteComment(id string, userId string) error
}
