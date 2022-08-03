package repository

import (
	//"database/sql"
	//"log"
	//"todo/models"

	"github.com/jmoiron/sqlx"
)

type todoRepository struct {
	db *sqlx.DB
}

type TodoRepository interface {
}

func InitTodoRepository(db *sqlx.DB) TodoRepository {
	return &todoRepository{
		db,
	}
}
