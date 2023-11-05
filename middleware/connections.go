package middleware

import (
	"github.com/gin-gonic/gin"
	"spark/database"
	"spark/utils"
)

func SetupConnections(db database.ConnectionsPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(utils.DatabaseKey, db)
	}
}
