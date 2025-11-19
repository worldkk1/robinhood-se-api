package repositories

import (
	"github.com/worldkk1/robinhood-se-api/domain"
	"github.com/worldkk1/robinhood-se-api/internal/database"
	"github.com/worldkk1/robinhood-se-api/internal/database/models"
	"gorm.io/gorm"
)

type taskPostgresRepository struct {
	db database.Database
}

func NewTaskPostgresRepository(db database.Database) TaskRepository {
	return &taskPostgresRepository{db: db}
}

func (r *taskPostgresRepository) Create(input domain.Task) error {
	return r.db.GetDb().Transaction(func(tx *gorm.DB) error {
		taskModel := models.TaskModel{
			Title:       input.Title,
			Description: input.Description,
			UserID:      input.UserID,
		}
		if err := tx.Create(&taskModel).Error; err != nil {
			return err
		}
		taskLog := models.TaskLogModel{
			TaskID:      taskModel.ID,
			Title:       &taskModel.Title,
			Description: &taskModel.Description,
			Status:      &taskModel.Status,
		}
		if err := tx.Create(&taskLog).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *taskPostgresRepository) Find(where string, params ...any) ([]domain.Task, error) {
	var taskModels []models.TaskModel
	if err := r.db.GetDb().Where(where, params...).Find(&taskModels).Error; err != nil {
		return nil, err
	}

	var result []domain.Task
	for _, m := range taskModels {
		result = append(result, domain.Task{
			ID:          m.ID,
			Title:       m.Title,
			Description: m.Description,
			UserID:      m.UserID,
			Status:      domain.TaskStatus(m.Status),
			ArchivedAt:  m.ArchivedAt,
			CreatedAt:   m.CreatedAt,
			UpdatedAt:   m.UpdatedAt,
		})
	}

	return result, nil
}

func (r *taskPostgresRepository) FindOneByID(id string) (*domain.Task, error) {
	var taskModel models.TaskModel
	if err := r.db.GetDb().Model(&models.TaskModel{}).Preload("User").Preload("TaskLogs").Where("id = ?", id).First(&taskModel).Error; err != nil {
		return nil, err
	}
	var taskLogs []domain.TaskLog
	for _, log := range taskModel.TaskLogs {
		var status domain.TaskStatus
		if log.Status != nil {
			status = domain.TaskStatus(*log.Status)
		}
		taskLogs = append(taskLogs, domain.TaskLog{
			ID:          log.ID,
			TaskID:      log.TaskID,
			Title:       log.Title,
			Description: log.Description,
			Status:      &status,
			CreatedAt:   log.CreatedAt,
			UpdatedAt:   log.UpdatedAt,
		})
	}
	result := domain.Task{
		ID:          taskModel.ID,
		Title:       taskModel.Title,
		Description: taskModel.Description,
		UserID:      taskModel.UserID,
		Status:      domain.TaskStatus(taskModel.Status),
		ArchivedAt:  taskModel.ArchivedAt,
		CreatedAt:   taskModel.CreatedAt,
		UpdatedAt:   taskModel.UpdatedAt,
		User:        domain.User(taskModel.User),
		TaskLog:     taskLogs,
	}

	return &result, nil
}

func (r *taskPostgresRepository) Update(id string, input domain.Task) error {
	return r.db.GetDb().Transaction(func(tx *gorm.DB) error {
		var taskModel models.TaskModel
		if err := tx.Where("id = ?", id).First(&taskModel).Error; err != nil {
			return err
		}

		taskLog := models.TaskLogModel{
			TaskID: taskModel.ID,
		}
		if input.Title != "" {
			taskModel.Title = input.Title
			taskLog.Title = &input.Title
		}
		if input.Description != "" {
			taskModel.Description = input.Description
			taskLog.Description = &input.Description
		}
		if input.Status != "" {
			status := models.TaskStatusEnum(input.Status)
			taskModel.Status = status
			taskLog.Status = &status
		}
		if input.ArchivedAt != nil {
			taskModel.ArchivedAt = input.ArchivedAt
		}

		if err := tx.Save(&taskModel).Error; err != nil {
			return err
		}
		if err := tx.Create(&taskLog).Error; err != nil {
			return err
		}

		return nil
	})
}
