package repository

import (
	"database/sql"
	"fmt"
	"github.com/jinseoko/real-life-go-api/common"
	"github.com/jinseoko/real-life-go-api/model"
)

type TodoRepository interface {
	WithTrx(trx *sql.Tx) TodoRepository
	Insert(title string, content string) (int64, error)
	FindOne(id int64) (*model.TodoModel, error)
	FindAll() ([]model.TodoModel, error)
}

type todoRepository struct {
	db common.QueryAble
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return todoRepository{db: db}
}

func (todoRepo todoRepository) WithTrx(trx *sql.Tx) TodoRepository {
	todoRepo.db = trx

	return todoRepo
}

func (todoRepo todoRepository) Insert(title string, content string) (int64, error) {
	query := `INSERT INTO todo (title, content)
		VALUES (?, ?)
	`

	result, err := todoRepo.db.Exec(query, title, content)

	if err != nil {
		return 0, fmt.Errorf("sql error on todoRepository.insert : %v", err)
	}

	return result.LastInsertId()
}

func (todoRepo todoRepository) FindOne(id int64) (*model.TodoModel, error) {
	var tm model.TodoModel

	query := `SELECT id,
					 title,
					 content,
					 created_at,
					 updated_at
			  FROM todo
			  WHERE id = ?`

	if err := todoRepo.db.QueryRow(query, id).Scan(&tm.Id, &tm.Title, &tm.Content, &tm.CreatedAt, &tm.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &tm, nil
}

func (todoRepo todoRepository) FindAll() ([]model.TodoModel, error) {
	var todos []model.TodoModel

	query := `
		SELECT id, title, content, created_at, updated_at
		FROM todo
	`

	rows, err := todoRepo.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var todo model.TodoModel

		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Content, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
