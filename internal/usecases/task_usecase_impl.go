package usecases

import (
	"time"

	"github.com/worldkk1/robinhood-se-api/domain"
	"github.com/worldkk1/robinhood-se-api/internal/repositories"
)

type taskUsecaseImpl struct {
	taskRepository repositories.TaskRepository
}

func NewTaskUsecaseImpl(taskRepository repositories.TaskRepository) TaskUsecase {
	return &taskUsecaseImpl{
		taskRepository: taskRepository,
	}
}

func (u *taskUsecaseImpl) CreateTask(input domain.Task) error {
	err := u.taskRepository.Create(input)
	if err != nil {
		return err
	}

	return nil
}

func (u *taskUsecaseImpl) GetTaskList(pagination Pagination) PaginationData[[]TaskList] {
	offset := pagination.Offset
	limit := pagination.Limit

	data, err := u.taskRepository.Find(repositories.FindOption{
		Where:  "archived_at IS NULL",
		Order:  "created_at asc",
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return PaginationData[[]TaskList]{Data: []TaskList{}, Total: 0}
	}

	tasks := data.Data
	var taskList []TaskList
	for _, task := range tasks {
		taskList = append(taskList, TaskList{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			UserID:      task.UserID,
			Status:      task.Status,
			ArchivedAt:  task.ArchivedAt,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
			User: TaskUserDetail{
				ID:    task.User.ID,
				Name:  task.User.Name,
				Email: task.User.Email,
			},
		})
	}

	return PaginationData[[]TaskList]{Data: taskList, Total: data.Total}
}

func (u *taskUsecaseImpl) GetTaskDetail(id string) *TaskDetail {
	task, err := u.taskRepository.FindOneByID(id)
	if err != nil {
		return nil
	}

	taskDetail := TaskDetail{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UserID:      task.UserID,
		ArchivedAt:  task.ArchivedAt,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		TaskLogs:    task.TaskLog,
		User: TaskUserDetail{
			ID:    task.User.ID,
			Name:  task.User.Name,
			Email: task.User.Email,
		},
	}

	return &taskDetail
}

func (u *taskUsecaseImpl) EditTask(id string, input domain.Task, updatedBy string) error {
	err := u.taskRepository.Update(id, input, updatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (u *taskUsecaseImpl) ArchiveTask(id string, updatedBy string) error {
	now := time.Now()
	err := u.taskRepository.Update(id, domain.Task{
		ArchivedAt: &now,
	}, updatedBy)
	if err != nil {
		return err
	}

	return nil
}
