package handlers

import (
	apiUtils "api-gateway/internal/utils"
	pb "api-gateway/shared/protobufs/document"
	"api-gateway/shared/utils"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetDocumentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileID := vars["id"]

	grpcClient := apiUtils.GetGRPCClient()

	stream, err := grpcClient.GetFileByID(r.Context(), &pb.GetFileRequest{FileId: fileID})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", `attachment; filename="file.pdf"`)

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		w.Write(res.FileChunk)
	}
}

func GetDocumentsHandler(w http.ResponseWriter, r *http.Request) {
	grpcClient := apiUtils.GetGRPCClient()

	resp, err := grpcClient.GetFileIDs(r.Context(), &pb.GetIDsRequest{})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		log.Printf("gRPC call failed: %v", err)
		return
	}

	if err = utils.WriteJSON(w, http.StatusOK, resp); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

func CreateDocumentHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 5<<20)

	if err := r.ParseMultipartForm(5 << 20); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to parse multiformat"))
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to retrieve file"))
		return
	}
	defer file.Close()

	grpcClient := apiUtils.GetGRPCClient() // Get the gRPC client instance
	stream, err := grpcClient.StoreFile(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Send the document metadata first
	if err := stream.Send(&pb.StoreFileRequest{
		Data: &pb.StoreFileRequest_Info{
			Info: &pb.DocumentInfo{
				Filename: fileHeader.Filename,
			},
		},
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to send document info"))
		return
	}

	// Send the document data in chunks
	buffer := make([]byte, 32*1024) // 32KB buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to read document"))
			return
		}

		if err := stream.Send(&pb.StoreFileRequest{
			Data: &pb.StoreFileRequest_FileChunk{
				FileChunk: buffer[:n],
			},
		}); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to send document chunk: %v", err))
			log.Fatal("cannot send chunk to server: ", err, stream.RecvMsg(nil))
			return
		}
	}

	// Close the stream and receive the response
	response, err := stream.CloseAndRecv()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to retrieve file"))
		return
	}
	
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"msg": "Document created successfully", "id": response.FileId})
}
