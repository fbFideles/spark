package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"spark/config"
)

func main() {
	config.LoadConfig()
	r := gin.Default()

	if err := r.Run(":" + viper.GetString("port")); err != nil {
		log.Fatal(err)
	}
}
