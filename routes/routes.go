package routes

import (
	"github.com/gin-gonic/gin"
	"quiz-api/handlers"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/quiz", handlers.CreateQuiz)
	r.POST("/question", handlers.CreateQuestion)
	r.POST("/start", handlers.StartAttempt)
	r.POST("/submit", handlers.SubmitAnswers)

}