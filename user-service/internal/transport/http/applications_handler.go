package http

import (
	"net/http"
	"user-service/internal/utils"
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

