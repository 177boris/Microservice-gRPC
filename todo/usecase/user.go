package usecase

//"encoding/json"
//"log"
//"todo/models"
//"todo/repository"

type todoUsecase struct {
	todoRepository repository.TodoRepository
}

type TodoUsecase interface {
}

func InitUserUsecase(todoRepository repository.TodoRepository) TodoUsecase {
	return &todoUsecase{
		todoRepository,
	}
}
