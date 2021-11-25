package service

import (
	"database/sql"
	"github.com/jinseoko/real-life-go-api/repository"
)

type TodoService interface {
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(tr repository.TodoRepository) TodoService {
	return todoService{todoRepository: tr}
}

func (ts todoService) withTrx(trx *sql.Tx) todoService {
	ts.todoRepository = ts.todoRepository.WithTrx(trx)
	return ts
}


