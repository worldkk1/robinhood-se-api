package usecases

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"github.com/worldkk1/robinhood-se-api/domain"
	"github.com/worldkk1/robinhood-se-api/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type signTokenPayload struct {
	UserID string
	Role   string
}

type authUsecaseImpl struct {
	userRepository repositories.UserRepository
}

func NewAuthUsecaseImpl(userRepository repositories.UserRepository) AuthUseCase {
	return &authUsecaseImpl{
		userRepository: userRepository,
	}
}

func (u *authUsecaseImpl) Register(input domain.User) error {
	hashPassword, err := getHashPassword(input.Password)
	if err != nil {
		return err
	}
	input.Password = hashPassword
	return u.userRepository.Create(input)
}

func (u *authUsecaseImpl) Login(email string, password string) (*AuthToken, error) {
	foundUser, err := u.userRepository.FindOne(domain.User{
		Email: email,
	})
	loginError := "email or password incorrect"
	if err != nil {
		return nil, errors.New(loginError)
	}
	if !checkPassword(foundUser.Password, password) {
		return nil, errors.New(loginError)
	}
	payload := signTokenPayload{
		UserID: foundUser.ID,
		Role:   foundUser.RoleID,
	}
	accessToken, err := signToken(payload, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, err
	}
	refreshToken, err := signToken(payload, time.Now().AddDate(0, 0, 30))
	if err != nil {
		return nil, err
	}

	return &AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func signToken(payload signTokenPayload, expiredAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  payload.UserID,
		"role": payload.Role,
		"exp":  expiredAt.Unix(),
	})
	secretKey := viper.GetString("SECRET_KEY")
	signToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signToken, nil
}

func getHashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
