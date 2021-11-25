package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinseoko/real-life-go-api/common"
	"github.com/jinseoko/real-life-go-api/repository"
	"log"
	"net/http"
)

type TestStruct struct {
	num int
}

func NewTestStruct(num int) TestStruct{
	return TestStruct{num: num}
}

func (t TestStruct) update(num int) TestStruct  {
	t.num = num
	return t
}

func withTrxMiddleWare(db *sql.DB, o *sql.TxOptions) func(c *gin.Context) {
	return func(c *gin.Context) {
		trx, err := db.BeginTx(c, o)
		if err != nil {
			log.Println("begin transaction error")
		}

		defer func() {
			if r := recover(); r != nil {
				trx.Rollback()
				log.Println(r, "trx do rollback")
			}
		}()

		c.Set("trx", trx)

		c.Next()

		if c.Writer.Status() == http.StatusOK || c.Writer.Status() == http.StatusCreated {
			if err := trx.Commit(); err != nil {
				log.Printf("trx commit error: %s", err)
			}
		} else {
			_ = trx.Rollback()
		}
	}
}

func main() {
	router := gin.Default()
	router.Group("/")

	dbConfig := common.GetMysqlDbConfig("localhost:3307", "root", "adfaie83ma", "golang-test")
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	todoRepository := repository.NewTodoRepository(db)
	if err = db.Ping(); err != nil {
		log.Println(err.Error())
	}

	todoModels, err := todoRepository.FindAll()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(todoModels)

}
