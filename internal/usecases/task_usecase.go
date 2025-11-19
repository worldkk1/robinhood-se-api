package usecases

import "github.com/worldkk1/robinhood-se-api/domain"

type TaskUsecase interface {
	CreateTask(input domain.Task) error
	GetTaskList() []domain.Task
	GetTaskDetail(id string) *domain.Task
	EditTask(id string, input domain.Task) error
	ArchiveTask(id string) error
}
