package repositories

import "github.com/worldkk1/robinhood-se-api/domain"

type TaskRepository interface {
	Create(input domain.Task) error
	Find(where string, params ...any) ([]domain.Task, error)
	FindOneByID(id string) (*domain.Task, error)
	Update(id string, input domain.Task, updatedBy string) error
}
