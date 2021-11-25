package repository

import (
	"database/sql"
)

type UserRepository interface {
	withTrx(trx *sql.Tx) UserRepository

}
