package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"quiz-api/config"
	"quiz-api/models"
)

type StartRequest struct {
	QuizID   uint   `json:"quizID"`
	UserName string `json:"userName"`
}

func StartAttempt(c *gin.Context) {

	var req StartRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	attempt := models.Attempt{
		QuizID:    req.QuizID,
		UserName:  req.UserName,
		StartTime: time.Now(),
	}

	config.DB.Create(&attempt)

	c.JSON(200, attempt)
}