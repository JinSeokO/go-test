package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinseoko/real-life-go-api/model"
	"github.com/jinseoko/real-life-go-api/testdata"
	"log"
	"reflect"
	"testing"
)

func TestNewTodoRepository(t *testing.T) {
	type args struct {
		db *sql.DB
	}

	db := testdata.GetTestDB()
	defer db.Close()

	tests := []struct {
		name string
		args args
		want TodoRepository
	}{
		{name: "testcase 1", args: args{db: db}, want: NewTodoRepository(db)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTodoRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodoRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoRepository_FindAll(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db := testdata.GetTestDB()
	defer db.Close()

	begin, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	todoRepo := todoRepository{db: db}.WithTrx(begin)
	todoModes, err := todoRepo.FindAll(5, 0)
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		want    []model.TodoModel
		wantErr bool
	}{
		{name: "test case1", fields: fields{db: db}, want: todoModes, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todoRepo := todoRepository{
				db: tt.fields.db,
			}
			got, err := todoRepo.FindAll(5, 0)

			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoRepository_FindOne(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.TodoModel
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todoRepo := todoRepository{
				db: tt.fields.db,
			}
			got, err := todoRepo.FindOne(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindOne() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoRepository_Insert(t *testing.T) {

	type args struct {
		title   string
		content string
	}

	db := testdata.GetTestDB()
	trx, err := db.Begin()
	if err != nil {
		t.Error(err)
	}
	todoRepo := NewTodoRepository(db)
	todoRepoWithTrx := todoRepo.WithTrx(trx)

	tests := []struct {
		name   string
		args   args
		expect args
	}{
		{"testcase1", args{"test case title1", "test case content1"}, args{"test case title1", "test case content1"}},
		{"testcase1", args{"test case title2", "test case content2"}, args{"test case title2", "test case content2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := todoRepoWithTrx.Insert(tt.args.title, tt.args.content)
			todo, err := todoRepoWithTrx.FindOne(id)

			if err != nil {
				t.Error(err)
			}

			if todo.Id != id {
				t.Errorf("InsertedId = %d, but got %d", id, todo.Id)

			}
			if todo.Title != tt.expect.title || todo.Content != tt.expect.content {
				t.Errorf("Insert title: %s, content: %s\n but  got title: %s, content %s",
					todo.Title, todo.Content, tt.expect.title, tt.expect.content)
			}
		})
	}
}

func Test_todoRepository_WithTrx(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		trx *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   todoRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todoRepo := todoRepository{
				db: tt.fields.db,
			}
			if got := todoRepo.WithTrx(tt.args.trx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithTrx() = %v, want %v", got, tt.want)
			}
		})
	}
}
