package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinseoko/real-life-go-api/common"
	"log"
)

func main() {
	router := gin.Default()
	router.Group("/")

	dbConfig := common.GetMysqlDbConfig("localhost:3307", "root", "adfaie83ma", "golang-test")
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err = db.Ping(); err != nil {
		log.Println(err.Error())
	}
}
