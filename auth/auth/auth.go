package auth

import (
	"auth/usecase"

	"golang.org/x/net/context"
)

type Server struct {
	userUsecase usecase.UserUsecase
	UnimplementedAuthServiceServer
}

func InitServer(userUsecase usecase.UserUsecase) Server {
	return Server{
		userUsecase,
		UnimplementedAuthServiceServer{},
	}
}

func (s *Server) Register(ctx context.Context, request *RegisterRequest) (*RegisterResponse, error) {
	return &RegisterResponse{Success: true, Message: "Succeed to register"}, nil
}

func (s *Server) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	return &LoginResponse{Success: true, Message: "Login success", Token: ""}, nil
}

func (s *Server) ValidateToken(ctx context.Context, request *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	return &ValidateTokenResponse{Success: true, Message: ""}, nil
}
