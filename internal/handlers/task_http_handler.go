package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/worldkk1/robinhood-se-api/domain"
	"github.com/worldkk1/robinhood-se-api/internal/dto"
	"github.com/worldkk1/robinhood-se-api/internal/handlers/middleware"
	"github.com/worldkk1/robinhood-se-api/internal/usecases"
)

type taskHttpHandler struct {
	taskUsecase usecases.TaskUsecase
}

func NewTaskHttpHandler(taskUsecase usecases.TaskUsecase) *taskHttpHandler {
	return &taskHttpHandler{
		taskUsecase: taskUsecase,
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
	tasks := h.taskUsecase.GetTaskList()
	result, err := json.Marshal(tasks)
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
	err = h.taskUsecase.EditTask(taskId, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *taskHttpHandler) ArchiveTask(w http.ResponseWriter, r *http.Request) {
	taskId := r.PathValue("id")
	err := h.taskUsecase.ArchiveTask(taskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
