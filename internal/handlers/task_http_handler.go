package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/worldkk1/robinhood-se-api/domain"
	"github.com/worldkk1/robinhood-se-api/internal/dto"
	"github.com/worldkk1/robinhood-se-api/internal/handlers/middleware"
	"github.com/worldkk1/robinhood-se-api/internal/usecases"
)

type taskHttpHandler struct {
	taskUsecase    usecases.TaskUsecase
	commentUsecase usecases.CommentUsecase
}

func NewTaskHttpHandler(taskUsecase usecases.TaskUsecase, commentUsecase usecases.CommentUsecase) *taskHttpHandler {
	return &taskHttpHandler{
		taskUsecase:    taskUsecase,
		commentUsecase: commentUsecase,
	}
}

func (h *taskHttpHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middleware.ContextUserKey).(middleware.AuthUser)
	if !ok {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	var payload dto.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	input := domain.Task{
		Title:       payload.Title,
		Description: payload.Description,
		UserID:      user.UserId,
	}
	if err := h.taskUsecase.CreateTask(input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *taskHttpHandler) GetTaskList(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	offset, _ := strconv.Atoi(queryParams.Get("offset"))
	if offset <= 0 {
		offset = 0
	}
	limit, _ := strconv.Atoi(queryParams.Get("limit"))
	if limit <= 0 {
		limit = 10
	}

	data := h.taskUsecase.GetTaskList(usecases.Pagination{
		Offset: offset,
		Limit:  limit,
	})
	tasks := data.Data
	if tasks == nil {
		tasks = []usecases.TaskList{}
	}
	response := dto.GetTaskListResponse{
		Offset: offset,
		Limit:  limit,
		Total:  data.Total,
		Data:   tasks,
	}
	result, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (h *taskHttpHandler) GetTaskDetail(w http.ResponseWriter, r *http.Request) {
	taskId := r.PathValue("id")
	task := h.taskUsecase.GetTaskDetail(taskId)
	if task == nil {
		http.Error(w, "data not found", http.StatusNotFound)
		return
	}
	result, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (h *taskHttpHandler) EditTask(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middleware.ContextUserKey).(middleware.AuthUser)
	if !ok {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	taskId := r.PathValue("id")
	var payload dto.EditTaskRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	input := domain.Task{
		Title:       payload.Title,
		Description: payload.Description,
		Status:      domain.TaskStatus(payload.Status),
	}
	err = h.taskUsecase.EditTask(taskId, input, user.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *taskHttpHandler) ArchiveTask(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middleware.ContextUserKey).(middleware.AuthUser)
	if !ok {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	taskId := r.PathValue("id")
	err := h.taskUsecase.ArchiveTask(taskId, user.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *taskHttpHandler) CreateTaskComment(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middleware.ContextUserKey).(middleware.AuthUser)
	if !ok {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	taskId := r.PathValue("id")
	var payload dto.CreateTaskCommentRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	input := domain.Comment{
		Content: payload.Content,
		UserID:  user.UserId,
		TaskID:  taskId,
	}
	err = h.commentUsecase.CreateComment(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *taskHttpHandler) GetTaskComments(w http.ResponseWriter, r *http.Request) {
	taskId := r.PathValue("id")
	comments := h.commentUsecase.GetTaskComments(taskId)
	if comments == nil {
		comments = []usecases.TaskComment{}
	}
	result, err := json.Marshal(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (h *taskHttpHandler) EditTaskComment(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middleware.ContextUserKey).(middleware.AuthUser)
	if !ok {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	commentId := r.PathValue("commentId")
	var payload dto.CreateTaskCommentRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.commentUsecase.EditComment(commentId, payload.Content, user.UserId)
	if err != nil {
		if err.Error() == "user_not_allow" {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *taskHttpHandler) DeleteTaskComment(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middleware.ContextUserKey).(middleware.AuthUser)
	if !ok {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	commentId := r.PathValue("commentId")
	err := h.commentUsecase.DeleteComment(commentId, user.UserId)
	if err != nil {
		if err.Error() == "user_not_allow" {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
