package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/worldkk1/robinhood-se-api/domain"
	"github.com/worldkk1/robinhood-se-api/internal/dto"
	"github.com/worldkk1/robinhood-se-api/internal/usecases"
)

type authHttpHandler struct {
	authUsecase usecases.AuthUseCase
}

func NewAuthHttpHandler(authUsecase usecases.AuthUseCase) *authHttpHandler {
	return &authHttpHandler{
		authUsecase: authUsecase,
	}
}

func (h *authHttpHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload dto.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	input := domain.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		RoleID:   payload.RoleID,
	}
	if err := h.authUsecase.Register(input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *authHttpHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	authToken, err := h.authUsecase.Login(payload.Email, payload.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := dto.LoginResponse{
		AccessToken:  authToken.AccessToken,
		RefreshToken: authToken.RefreshToken,
	}
	result, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
