package middleware

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"spark/utils"
)

func SetupConnections(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(utils.DatabaseKey, db)
	}
}
