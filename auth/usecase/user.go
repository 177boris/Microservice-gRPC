package usecase

import (
	"auth/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

type UserUsecase interface {
	Register(username, password string) (bool, error)
	Login(username, password string) (string, error)
	ValidateToken(token string) (string, error)
}

func InitUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository,
	}
}

func (userUsecase *userUsecase) Register(username, password string) (bool, error) {
	return true, nil
}

func (userUsecase *userUsecase) Login(username, password string) (string, error) {
	return "", nil
}

func (userUsecase *userUsecase) ValidateToken(token string) (string, error) {
	return "", nil
}
