package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"user-service/internal/utils"
	"user-service/types"

	"github.com/gorilla/mux"
	// "user-service/types"
)

func (h *Handler) ListApplications(w http.ResponseWriter, r *http.Request) {
	applications, err := h.applService.ListApplications(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, applications); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

func (h *Handler) GetApplicationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	applicationID, err := strconv.Atoi(idStr)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	application, err := h.applService.GetApplicationByID(r.Context(), applicationID)

	if err != nil {
		if err.Error() == "application not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, application); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

func (h *Handler) ListApplicationsByFarmerID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	fmt.Print(idStr)

	farmerID, err := strconv.Atoi(idStr)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	applications, err := h.applService.ListApplicationsByFarmerID(r.Context(), farmerID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, applications); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

func (h *Handler) UpdateApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	applicationID, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var payload types.ApplicationUpdatePayload
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.applService.UpdateApplication(r.Context(), applicationID, payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, "application updated successfully")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}
