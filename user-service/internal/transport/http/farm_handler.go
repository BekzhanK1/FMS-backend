package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-service/types"
)

func (h *Handler) CreateFarmHandler(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateFarmPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if payload.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	err = h.farmService.CreateFarm(r.Context(), payload.Name, payload.Address, payload.GeoLoc, payload.Size, payload.CropTypes, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "Farm created successfully")
}
