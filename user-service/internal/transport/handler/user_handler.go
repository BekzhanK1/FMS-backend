package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"user-service/internal/service"
	"user-service/internal/utils"
	"user-service/types"

	"github.com/go-playground/validator/v10"
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
	router.HandleFunc("/login", h.LoginHandler).Methods(http.MethodPost)
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
	var payload types.CreateUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	_, err := h.service.CreateUser(payload.Email, payload.Username, payload.Phone, payload.Password, false, payload.Role, payload.ProfilePicture)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusCreated, map[string]string{"msg": "User created successfully"}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
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
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id: %v", err))
		return
	}

	var payload types.UpdateUserPayload
	if err = utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err = utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err = h.service.UpdateUser(id, payload.Username, payload.Phone, payload.ProfilePicture, payload.IsActive)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err = utils.WriteJSON(w, http.StatusCreated, map[string]string{"msg": fmt.Sprintf("User with id %d updated successfully", id)}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	id, _ := strconv.Atoi(idStr)

	if err := h.service.DeleteUser(id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, map[string]string{"msg": "Deleted successfully"}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}


