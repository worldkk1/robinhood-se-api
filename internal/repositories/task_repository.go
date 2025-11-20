package repositories

import "github.com/worldkk1/robinhood-se-api/domain"

type FindOption struct {
	Where       string
	WhereParams []any
	Order       string
	Offset      int
	Limit       int
}

type PaginationData[T any] struct {
	Total int64
	Data  T
}

type TaskRepository interface {
	Create(input domain.Task) error
	Find(option FindOption) (PaginationData[[]domain.Task], error)
	FindOneByID(id string) (*domain.Task, error)
	Update(id string, input domain.Task, updatedBy string) error
}
