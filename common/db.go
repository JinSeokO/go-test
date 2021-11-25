package common

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type QueryAble interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func GetMysqlDbConfig(addr string, user string, password string, dbName string) *mysql.Config {
	dbConfig := mysql.NewConfig()
	dbConfig.Addr = addr
	dbConfig.User = user
	dbConfig.Passwd = password
	dbConfig.DBName = dbName
	dbConfig.Net = "tcp"
	dbConfig.ParseTime = true
	return dbConfig
}