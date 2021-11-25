package middleware

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func WithTrxMiddleWare(db *sql.DB, o *sql.TxOptions) func(c *gin.Context) {
	return func(c *gin.Context) {
		trx, err := db.BeginTx(c, o)
		if err != nil {
			log.Println("begin transaction error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
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
