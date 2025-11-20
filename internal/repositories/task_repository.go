package repositories

import "github.com/worldkk1/robinhood-se-api/domain"

type FindOption struct {
	Where       string
	WhereParams []any
	Order       string
}

type TaskRepository interface {
	Create(input domain.Task) error
	Find(option FindOption) ([]domain.Task, error)
	FindOneByID(id string) (*domain.Task, error)
	Update(id string, input domain.Task, updatedBy string) error
}
