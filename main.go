package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"spark/config"
	"spark/database"
	"spark/handler/pessoas"
	"spark/middleware"
)

func main() {
	config.LoadConfig()

	db, err := database.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	v1 := r.Group("v1")
	v1.Use(
		middleware.SetupConnections(db),
	)

	pessoas.Router(v1)

	if err := r.Run(":" + viper.GetString("port")); err != nil {
		log.Fatal(err)
	}
}
