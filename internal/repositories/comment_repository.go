package repositories

import "github.com/worldkk1/robinhood-se-api/domain"

type CommentRepository interface {
	Create(input domain.Comment) error
	Find(option FindOption) ([]domain.Comment, error)
	FindOneByID(id string) (*domain.Comment, error)
	Update(id string, input domain.Comment) error
	Delete(id string) error
}
