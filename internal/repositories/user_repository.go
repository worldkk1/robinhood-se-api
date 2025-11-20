package repositories

import "github.com/worldkk1/robinhood-se-api/domain"

type UserRepository interface {
	Create(input domain.User) error
	FindOne(where domain.User) (*domain.User, error)
}
