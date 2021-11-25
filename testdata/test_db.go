package testdata

import (
	"database/sql"
	"github.com/jinseoko/real-life-go-api/common"
)
import _ "github.com/go-sql-driver/mysql"
const DbHost = "localhost:3307"
const DbUser = "root"
const DbPassword = "adfaie83ma"
const DbDatabaseName = "golang-test"

func GetTestDB() *sql.DB {
	mysqlConfig := common.GetMysqlDbConfig(DbHost, DbUser, DbPassword, DbDatabaseName)
	open, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	return open
}