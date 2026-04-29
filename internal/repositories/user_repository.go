package repositories

import "github.com/worldkk1/work-ticket-based-api/domain"

type UserRepository interface {
	Create(input domain.User) error
	FindOne(where domain.User) (*domain.User, error)
}
