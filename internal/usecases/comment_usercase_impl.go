package usecases

import (
	"errors"

	"github.com/worldkk1/robinhood-se-api/domain"
	"github.com/worldkk1/robinhood-se-api/internal/repositories"
)

type commentUsecaseImpl struct {
	commentRepository repositories.CommentRepository
}

func NewCommentUsecaseImpl(commentRepository repositories.CommentRepository) CommentUsecase {
	return &commentUsecaseImpl{
		commentRepository: commentRepository,
	}
}

func (u *commentUsecaseImpl) CreateComment(input domain.Comment) error {
	err := u.commentRepository.Create(input)

	return err
}

func (u *commentUsecaseImpl) GetTaskComments(taskId string) []domain.Comment {
	comments, err := u.commentRepository.Find("task_id = ?", taskId)
	if err != nil {
		return []domain.Comment{}
	}

	return comments
}

func (u *commentUsecaseImpl) EditComment(id string, content string, userId string) error {
	comment, err := u.commentRepository.FindOneByID(id)
	if err != nil {
		return err
	}
	if comment.UserID != userId {
		return errors.New("user_not_allow")
	}

	return u.commentRepository.Update(id, domain.Comment{
		Content: content,
	})
}

func (u *commentUsecaseImpl) DeleteComment(id string, userId string) error {
	comment, err := u.commentRepository.FindOneByID(id)
	if err != nil {
		return err
	}
	if comment.UserID != userId {
		return errors.New("user_not_allow")
	}
	return u.commentRepository.Delete(id)
}
