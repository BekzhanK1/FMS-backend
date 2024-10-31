package http

import (
	"fmt"
	"net/http"
	"strconv"
	"user-service/internal/utils"

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
		utils.WriteError(w, http.StatusInternalServerError, err)
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
