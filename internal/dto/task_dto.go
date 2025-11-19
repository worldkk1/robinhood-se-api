package dto

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
