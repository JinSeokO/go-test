package service

import (
	"database/sql"
	"github.com/jinseoko/real-life-go-api/model"
	"github.com/jinseoko/real-life-go-api/repository"
)

type TodoService interface {
	WithTrx(trx *sql.Tx) TodoService
	GetTodos(limit, offset int) ([]model.TodoModel, error)
	CreateTodo(title, content string) (int64, error)
	GetTodoById(int64) (*model.TodoModel, error)
	GetTodoByTitleOrLikeComment(title string, content string) (*model.TodoModel, error)
}

func (ts todoService) GetTodoByTitleOrLikeComment(title string, content string) (*model.TodoModel, error) {
	return ts.todoRepository.FindByTitleOrLikeContent(title, content)
}

func NewTodoService(tr repository.TodoRepository) TodoService {
	return todoService{todoRepository: tr}
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func (ts todoService) WithTrx(trx *sql.Tx) TodoService {
	ts.todoRepository = ts.todoRepository.WithTrx(trx)
	return ts
}

func (ts todoService) GetTodos(limit int, offset int) ([]model.TodoModel, error) {
	return ts.todoRepository.FindAll(limit, offset)
}

func (ts todoService) CreateTodo(title, content string) (int64, error) {
	return ts.todoRepository.Insert(title, content)
}

func (ts todoService) GetTodoById(id int64) (*model.TodoModel, error) {
	return ts.todoRepository.FindOne(id)
}
