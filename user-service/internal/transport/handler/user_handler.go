// internal/handler/user_handler.go

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-service/internal/models"
	"user-service/internal/service"

	"github.com/gorilla/mux"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
// CreateUserHandler handles user creation
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
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

	user, err := service.CreateUser(
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

// GetUserHandler retrieves a user by ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from the URL path using mux.Vars
	vars := mux.Vars(r)
	idStr := vars["id"] // The "id" here matches the name in the route

	// Convert the ID from string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Call the service to get the user
	user, err := service.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if the user was found
	if user == nil {
		http.NotFound(w, r)
		return
	}

	// Return the user as JSON
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateUserHandler updates a user
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
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

	err := service.UpdateUser(
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

// DeleteUserHandler removes a user by ID
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = service.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
