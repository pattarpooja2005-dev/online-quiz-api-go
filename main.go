package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"quiz-api/config"
	"quiz-api/routes"
)

func main() {
	println("SERVER STARTING...")

	// connect database
	config.ConnectDatabase()

	// create router
	r := gin.Default()

	// enable CORS
	r.Use(cors.Default())

	// ⭐ load HTML templates
	r.LoadHTMLGlob("templates/*")

	// ⭐ serve static files (CSS, JS)
	r.Static("/static", "./static")

	// ⭐ homepage (UI)
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// ⭐ quiz page
	r.GET("/quiz", func(c *gin.Context) {
		c.HTML(200, "quiz.html", nil)
	})

	// ⭐ result page
	r.GET("/result", func(c *gin.Context) {
		c.HTML(200, "result.html", nil)
	})

	// ⭐ API routes (your existing endpoints)
	routes.SetupRoutes(r)

	// run server
	r.Run(":8080")
}