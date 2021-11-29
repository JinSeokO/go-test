package service

import (
	"database/sql"
	"github.com/jinseoko/real-life-go-api/repository"
	"github.com/jinseoko/real-life-go-api/testdata"
	"reflect"
	"testing"
)

var db *sql.DB
var ts TodoService

func init() {
	db = testdata.GetTestDB()
	todoRepository := repository.NewTodoRepository(db)
	ts = NewTodoService(todoRepository)
}

func setUp() func() {
	trx, err := db.Begin()
	originTodoService := ts
	if err != nil {
		panic(err)
	}
	ts = ts.WithTrx(trx)

	// setDown
	return func() {
		ts = originTodoService
		_ = trx.Rollback()
	}
}

func TestNewTodoService(t *testing.T) {
	type args struct {
		tr repository.TodoRepository
	}

	db := testdata.GetTestDB()
	todoRepository := repository.NewTodoRepository(db)

	tests := []struct {
		name string
		args args
		want TodoService
	}{
		{"Either todoRepository same or not", args{todoRepository}, NewTodoService(todoRepository)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer setUp()()
			if got := NewTodoService(tt.args.tr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodoService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoService_CreateTodo_And_FindTodo(t *testing.T) {
	type args struct {
		title   string
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    args
		wantErr bool
	}{
		{"Either create todo and find todo equal", args{"test title1",
			"test content1"}, args{"test title1", "test content1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer setUp()()
			got, err := ts.CreateTodo(tt.args.title, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
			todoModel, err := ts.GetTodoById(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTodoById() error = %v, wantErr %v", err, tt.wantErr)
			}

			if todoModel.Title != tt.want.title {
				t.Errorf("GetTodoById().title got = %v, want %v", todoModel.Title, tt.want.title)
			}
			if todoModel.Content != tt.want.content {
				t.Errorf("GetTodoById().Conetent got = %v, want %v", todoModel.Content, tt.want.content)
			}
		})
	}
}
