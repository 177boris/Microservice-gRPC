package repository

import (
	"auth/models"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

type UserRepository interface {
	CreateUser(username string, password string) (bool, error)
	GetUserByUsername(username string) (models.User, error)
}

func InitUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db,
	}
}

func (userRepository *userRepository) CreateUser(username string, password string) (bool, error) {
	return true, nil
}

func (userRepository *userRepository) GetUserByUsername(username string) (models.User, error) {
	return models.User{}, nil
}
