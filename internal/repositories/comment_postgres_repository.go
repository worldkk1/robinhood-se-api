package repositories

import (
	"github.com/worldkk1/robinhood-se-api/domain"
	"github.com/worldkk1/robinhood-se-api/internal/database"
	"github.com/worldkk1/robinhood-se-api/internal/database/models"
)

type commentPostgresRepository struct {
	db database.Database
}

func NewCommentPostgresRepository(db database.Database) CommentRepository {
	return &commentPostgresRepository{db: db}
}

func (r *commentPostgresRepository) Create(input domain.Comment) error {
	commentModel := models.CommentModel{
		Content: input.Content,
		UserID:  input.UserID,
		TaskID:  input.TaskID,
	}
	if err := r.db.GetDb().Create(&commentModel).Error; err != nil {
		return err
	}

	return nil
}

func (r *commentPostgresRepository) Find(where string, params ...any) ([]domain.Comment, error) {
	var commentModels []models.CommentModel
	if err := r.db.GetDb().Where(where, params...).Find(&commentModels).Error; err != nil {
		return nil, err
	}

	var result []domain.Comment
	for _, m := range commentModels {
		result = append(result, domain.Comment{
			ID:        m.ID,
			Content:   m.Content,
			UserID:    m.UserID,
			TaskID:    m.TaskID,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		})
	}

	return result, nil
}

func (r *commentPostgresRepository) Update(id string, input domain.Comment) error {
	err := r.db.GetDb().Model(&models.CommentModel{}).Where("id = ?", id).Update("content", input.Content).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *commentPostgresRepository) Delete(id string) error {
	return r.db.GetDb().Delete(&models.CommentModel{ID: id}).Error
}

func (r *commentPostgresRepository) FindOneByID(id string) (*domain.Comment, error) {
	var commentModel models.CommentModel
	if err := r.db.GetDb().Where("id = ?", id).First(&commentModel).Error; err != nil {
		return nil, err
	}

	result := domain.Comment{
		ID:        commentModel.ID,
		Content:   commentModel.Content,
		UserID:    commentModel.UserID,
		TaskID:    commentModel.TaskID,
		CreatedAt: commentModel.CreatedAt,
		UpdatedAt: commentModel.UpdatedAt,
	}

	return &result, nil
}
