package utils

import (
	pb "api-gateway/shared/protobufs/document"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var grpcClient pb.DocumentServiceClient

func InitGRPCClient(address string) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	grpcClient = pb.NewDocumentServiceClient(conn)

	fmt.Println("Connected to DocumentService")
}

func GetGRPCClient() pb.DocumentServiceClient {
	return grpcClient
}
