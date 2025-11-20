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

func (u *commentUsecaseImpl) GetTaskComments(taskId string) []TaskComment {
	comments, err := u.commentRepository.Find(repositories.FindOption{
		Where:       "task_id = ?",
		WhereParams: []any{taskId},
		Order:       "created_at desc",
	})
	if err != nil {
		return []TaskComment{}
	}

	var commentList []TaskComment
	for _, comment := range comments {
		commentList = append(commentList, TaskComment{
			ID:        comment.ID,
			Content:   comment.Content,
			UserID:    comment.UserID,
			TaskID:    comment.TaskID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			User: CommentUser{
				ID:    comment.User.ID,
				Name:  comment.User.Name,
				Email: comment.User.Email,
			},
		})
	}

	return commentList
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
