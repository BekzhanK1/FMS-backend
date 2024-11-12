package grpc

import (
	"context"

	"user-service/internal/models"
	authService "user-service/internal/service/auth"
	userService "user-service/internal/service/user"
	pb "user-service/shared/protobufs/user-service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	userService *userService.Service
	authService *authService.Service
}

func NewServer(grpcServer *grpc.Server, userService *userService.Service, authService *authService.Service) {
	usrGrpcServer := &Server{userService: userService, authService: authService}
	pb.RegisterUserServiceServer(grpcServer, usrGrpcServer)
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	encryptedEmail, err := s.userService.CreateUser(req.Email, req.Username, req.FirstName, req.LastName, req.PhoneNumber, req.Password, false, models.Role(req.GetRole().String()), req.ProfilePicture)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	return &pb.CreateUserResponse{Message: "Created Successfully", EncryptedEmail: encryptedEmail}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	tokens, err := s.authService.Login(req.Email, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to login: %v", err)
	}

	return &pb.LoginResponse{AccessToken: tokens.AccessToken, RefreshToken: tokens.RefreshToken}, nil
}

func (s *Server) ActivateUser(ctx context.Context, req *pb.ActivateUserRequest) (*pb.ActivateUserResponse, error) {
	err := s.userService.ActivateUser(req.EncryptedEmail, req.Otp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to activate user: %v", err)
	}

	return &pb.ActivateUserResponse{Message: "User activated successfully"}, nil
}
