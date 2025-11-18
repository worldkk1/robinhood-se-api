package repositories

import (
	"github.com/worldkk1/robinhood-se-api/domain"
	"github.com/worldkk1/robinhood-se-api/internal/database"
	"github.com/worldkk1/robinhood-se-api/internal/database/models"
)

type userPostgresRepository struct {
	db database.Database
}

func NewUserPostgresRepository(db database.Database) UserRepository {
	return &userPostgresRepository{db: db}
}

func (r *userPostgresRepository) Create(input domain.User) error {
	userModel := models.UserModel{
		ID:       "",
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		RoleID:   input.RoleID,
	}
	if err := r.db.GetDb().Create(&userModel).Error; err != nil {
		return err
	}

	return nil
}

func (r *userPostgresRepository) FindOne(where domain.User) (*domain.User, error) {
	var userModel models.UserModel
	if err := r.db.GetDb().Where(&models.UserModel{
		ID:    where.ID,
		Name:  where.Name,
		Email: where.Email,
	}).First(&userModel).Error; err != nil {
		return nil, err
	}
	user := domain.User{
		ID:        userModel.ID,
		Name:      userModel.Name,
		Email:     userModel.Email,
		Password:  userModel.Password,
		RoleID:    userModel.RoleID,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}

	return &user, nil
}
