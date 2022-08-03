package todo

import (
	context "context"
	//"fmt"
	//"log"
	//"todo/models"
	"todo/usecase"
)

type Server struct {
	todoUsecase usecase.TodoUsecase
	UnimplementedTodoServiceServer
}

func InitServer(todoUsecase usecase.TodoUsecase) Server {
	return Server{
		todoUsecase,
		UnimplementedTodoServiceServer{},
	}
}

func (s *Server) CreateTodo(ctx context.Context, request *CreateTodoRequest) (*CreateTodoResponse, error) {
	return &CreateTodoResponse{Success: true, Message: "Success to create todo"}, nil
}

func (s *Server) GetTodos(ctx context.Context, request *GetTodosRequest) (*GetTodosResponse, error) {
	return &GetTodosResponse{Success: true, Message: "Success to get todos", Data: ""}, nil
}

func (s *Server) UpdateTodo(ctx context.Context, request *UpdateTodoRequest) (*UpdateTodoResponse, error) {
	return &UpdateTodoResponse{Success: true, Message: "Success to update todo"}, nil
}

func (s *Server) DeleteTodo(ctx context.Context, request *DeleteTodoRequest) (*DeleteTodoResponse, error) {
	return &DeleteTodoResponse{Success: true, Message: "Success to delete todo"}, nil
}
