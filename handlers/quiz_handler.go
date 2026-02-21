package handlers

import (
	"github.com/gin-gonic/gin"
	"quiz-api/config"
	"quiz-api/models"
)

func CreateQuiz(c *gin.Context) {

	var quiz models.Quiz

	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&quiz)
	c.JSON(200, quiz)
}