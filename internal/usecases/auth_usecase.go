package usecases

import (
	"github.com/worldkk1/robinhood-se-api/domain"
)

type AuthToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthUseCase interface {
	Register(input domain.User) error
	Login(email string, password string) (*AuthToken, error)
}
