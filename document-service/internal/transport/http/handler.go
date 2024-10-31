package http

import (
	documentService "document-service/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	documentService documentService.Service
}

func NewHanlder(documentService documentService.Service) *Handler {
	return &Handler{
		documentService,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("", h.CreateDocumentHandler).Methods(http.MethodPost)
	router.HandleFunc("", h.GetIDsHandler).Methods(http.MethodGet)
	router.HandleFunc("/{fileId}", h.GetDocumentHandler).Methods(http.MethodGet)
}
