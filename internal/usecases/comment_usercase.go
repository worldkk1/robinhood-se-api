package usecases

import "github.com/worldkk1/robinhood-se-api/domain"

type CommentUsecase interface {
	CreateComment(input domain.Comment) error
	GetTaskComments(taskId string) []domain.Comment
	EditComment(id string, content string, userId string) error
	DeleteComment(id string, userId string) error
}
