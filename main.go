package main

import (
	"github.com/gin-gonic/gin"
	"quiz-api/config"
	"quiz-api/routes"
)

func main() {

	config.ConnectDatabase()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}