package http

import (
	"document-service/shared/utils"
	"document-service/types"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetIDsHandler(w http.ResponseWriter, r *http.Request) {
	ids, err := h.documentService.GetFileIDs(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err = utils.WriteJSON(w, http.StatusOK, ids); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}


func (h *Handler) GetDocumentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["fileId"]

	fileID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid file ID: %v", err))
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.pdf\"", idStr))


	err = h.documentService.GetFileByID(r.Context(), fileID, w)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to retrieve file: %v", err))
		return
	}
}


func (h *Handler) CreateDocumentHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 5<<20)
	// Parse the multipart form data
	if err := r.ParseMultipartForm(5 << 20); err != nil { // 10 MB max file size
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to parse multipart form: %v", err))
		return
	}

	// Retrieve the file from the form data
	file, fileHeader, err := r.FormFile("file") // "file" should match the form field name
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to retrieve file: %v", err))
		return
	}
	defer file.Close()

	// Populate the payload with file information
	payload := types.CreateDocumentPayload{
		FileHeader: *fileHeader,
		File:       file,
	}

	// Call the service to store the file
	id, err := h.documentService.CreateFile(r.Context(), payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with success message and file ID
	if err := utils.WriteJSON(w, http.StatusCreated, map[string]string{"msg": "Document created successfully", "key": id}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
