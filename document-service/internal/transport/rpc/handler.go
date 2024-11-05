package rpc

import (
	"context"
	"document-service/internal/service"
	pb "document-service/shared/protobufs/document"
	"document-service/types"
	"io"
	"mime/multipart"
	"os"

	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type Server struct {
    pb.UnimplementedDocumentServiceServer
    documentService *service.Service
}

func NewServer(grpcServer *grpc.Server, documentService *service.Service) {
    docGrpcServer := &Server{documentService: documentService}
    pb.RegisterDocumentServiceServer(grpcServer, docGrpcServer)
}

// GetFileIDs handles the GetFileIDs RPC
func (s *Server) GetFileIDs(ctx context.Context, req *pb.GetIDsRequest) (*pb.GetIDsResponse, error) {
    ids, err := s.documentService.GetFileIDs(ctx)
    if err != nil {
        log.Printf("Error retrieving file IDs: %v", err)
        return nil, err
    }

    // Convert ObjectIDs to string
    var stringIDs []string
    for _, id := range ids {
        stringIDs = append(stringIDs, id.Hex())
    }

    return &pb.GetIDsResponse{FileIds: stringIDs}, nil
}

// GetFileByID handles the GetFileByID RPC
func (s *Server) GetFileByID(req *pb.GetFileRequest, stream pb.DocumentService_GetFileByIDServer) error {
    fileID, err := primitive.ObjectIDFromHex(req.FileId)
    if err != nil {
        log.Printf("Invalid file ID: %v", err)
        return err
    }

    pr, pw := io.Pipe()
    go func() {
        defer pw.Close()
        if err := s.documentService.GetFileByID(context.Background(), fileID, pw); err != nil {
            log.Printf("Error retrieving file: %v", err)
        }
    }()

    buffer := make([]byte, 1024*32) // 32KB buffer
    for {
        n, err := pr.Read(buffer)
        if err != nil && err != io.EOF {
            log.Printf("Error reading file: %v", err)
            return err
        }
        if n == 0 {
            break
        }
        stream.Send(&pb.GetFileResponse{FileChunk: buffer[:n]})
    }

    return nil
}

func (s *Server) StoreFile(stream pb.DocumentService_StoreFileServer) error {
    var filename string
    tmpFile, err := os.CreateTemp("", "document-*.tmp")
    if err != nil {
        log.Printf("Error creating temporary file: %v", err)
        return err
    }
    defer tmpFile.Close()

    for {
        req, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Printf("Error receiving file chunk: %v", err)
            return err
        }

        if filename == "" {
            filename = req.GetInfo().Filename
            log.Printf("Receiving file: %s", filename)
        }

        if _, err := tmpFile.Write(req.GetFileChunk()); err != nil {
            log.Printf("Error writing to temporary file: %v", err)
            return err
        }
    }

    if _, err := tmpFile.Seek(0, io.SeekStart); err != nil {
        log.Printf("Error rewinding temporary file: %v", err)
        return err
    }

    fileInfo, _ := tmpFile.Stat()

    fileHeader := &multipart.FileHeader{
        Filename: filename,
        Size: fileInfo.Size(),
    }

    payload := types.CreateDocumentPayload{
        FileHeader: *fileHeader,
        File:       tmpFile,
    }

    fileID, err := s.documentService.CreateFile(context.Background(), payload)
    if err != nil {
        log.Printf("Error storing file: %v", err)
        return err
    }

    return stream.SendAndClose(&pb.StoreFileResponse{FileId: fileID})
}
