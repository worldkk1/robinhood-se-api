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

func (u *taskUsecaseImpl) GetTaskList() []domain.Task {
	tasks, err := u.taskRepository.Find("archived_at IS NULL")
	if err != nil {
		return []domain.Task{}
	}

	return tasks
}

func (u *taskUsecaseImpl) GetTaskDetail(id string) *domain.Task {
	task, err := u.taskRepository.FindOneByID(id)
	if err != nil {
		return nil
	}

	return task
}

func (u *taskUsecaseImpl) EditTask(id string, input domain.Task) error {
	err := u.taskRepository.Update(id, input)
	if err != nil {
		return err
	}

	return nil
}

func (u *taskUsecaseImpl) ArchiveTask(id string) error {
	now := time.Now()
	err := u.taskRepository.Update(id, domain.Task{
		ArchivedAt: &now,
	})
	if err != nil {
		return err
	}

	return nil
}
