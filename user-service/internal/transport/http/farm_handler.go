package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-service/types"
)

func (h *Handler) CreateFarmHandler(w http.ResponseWriter, r *http.Request) {
	// Decode JSON body into CreateFarmPayload struct
	var payload types.CreateFarmPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Check required fields
	if payload.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	// Extract user ID from context (assuming you have a way to get user ID from the request context)
	userId, err := getUserIDFromContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Call the service layer to create the farm
	err = h.farmService.CreateFarm(userId, payload.Name, payload.Address, payload.GeoLoc, payload.Size, payload.CropTypes, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Success response
	fmt.Fprintln(w, "Farm created successfully")
}
