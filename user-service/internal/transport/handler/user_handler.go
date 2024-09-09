package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-service/internal/models"
	"user-service/internal/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service service.Service
}

func NewHanlder(service service.Service) *Handler {
	return &Handler{
		service,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("", h.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/{id}", h.GetUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/{id}", h.UpdateUserHandler).Methods(http.MethodPut)
	router.HandleFunc("/{id}", h.DeleteUserHandler).Methods(http.MethodDelete)
}


func (h *Handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Email           string        `json:"email"`
		Username        string        `json:"username"`
		Phone           string        `json:"phone"`
		PasswordHash    string        `json:"password_hash"`
		IsActive        bool          `json:"is_active"`
		Role            models.Role   `json:"role"`
		ProfilePicture  string        `json:"profile_picture"`
	}
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(
		userInput.Email,
		userInput.Username,
		userInput.Phone,
		userInput.PasswordHash,
		userInput.IsActive,
		userInput.Role,
		userInput.ProfilePicture,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		ID              int           `json:"id"`
		Email           string        `json:"email"`
		Username        string        `json:"username"`
		Phone           string        `json:"phone"`
		PasswordHash    string        `json:"password_hash"`
		IsActive        bool          `json:"is_active"`
		Role            models.Role   `json:"role"`
		ProfilePicture  string        `json:"profile_picture"`
	}
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.UpdateUser(
		userInput.ID,
		userInput.Email,
		userInput.Username,
		userInput.Phone,
		userInput.PasswordHash,
		userInput.IsActive,
		userInput.Role,
		userInput.ProfilePicture,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
