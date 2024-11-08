package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"user-service/shared/utils"
	"user-service/types"

	"github.com/gorilla/mux"
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

func (h *Handler) ListFarms(w http.ResponseWriter, r *http.Request) {
	farms, err := h.farmService.ListFarms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, farms); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

func (h *Handler) GetFarmByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid farm ID", http.StatusBadRequest)
		return
	}

	farm, err := h.farmService.GetFarmByID(id)
	if err != nil {
		if err.Error() == "farm not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, farm); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

func (h *Handler) ListFarmsByFarmerID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	farmerID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid farmer ID", http.StatusBadRequest)
		return
	}

	farms, err := h.farmService.ListFarmsByFarmerID(farmerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, farms); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}
