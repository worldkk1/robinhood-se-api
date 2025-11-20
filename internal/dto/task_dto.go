package dto

import "github.com/worldkk1/robinhood-se-api/internal/usecases"

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type EditTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CreateTaskCommentRequest struct {
	Content string `json:"content"`
}

type GetTaskListResponse struct {
	Offset int                 `json:"offset"`
	Limit  int                 `json:"limit"`
	Total  int64               `json:"total"`
	Data   []usecases.TaskList `json:"data"`
}
