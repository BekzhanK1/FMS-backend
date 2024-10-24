package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"user-service/internal/middleware"
	"user-service/internal/utils"
	"user-service/types"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

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

	encryptedEmail, err := h.userService.CreateUser(payload.Email, payload.Username, payload.Phone, payload.Password, false, payload.Role, payload.ProfilePicture)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusCreated, map[string]string{"msg": "User created successfully", "key": encryptedEmail}); err != nil {
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

	user, err := h.userService.GetUserById(id)
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

	err = h.userService.UpdateUser(id, payload.Username, payload.Phone, payload.ProfilePicture, payload.IsActive)

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

	if err := h.userService.DeleteUser(id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, map[string]string{"msg": "Deleted successfully"}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}



func (h *Handler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	
	user, err := h.userService.GetUserById(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err = utils.WriteJSON(w, http.StatusOK, user); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}


func getUserIDFromContext(r *http.Request) (int, error) {
	userIdstr, ok := r.Context().Value(middleware.UserKey).(string)
	if !ok {
		return 0, fmt.Errorf("unable to get user ID from context")
	}
	userId, err := strconv.Atoi(userIdstr)
	if err != nil {
		return 0, err
	}
	return userId, nil	
}

