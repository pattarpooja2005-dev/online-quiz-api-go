package handlers

import (
	"time"
	"github.com/gin-gonic/gin"
	"quiz-api/config"
	"quiz-api/models"
)

func StartAttempt(c *gin.Context) {

	var attempt models.Attempt

	if err := c.ShouldBindJSON(&attempt); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	attempt.StartTime = time.Now()

	config.DB.Create(&attempt)
	c.JSON(200, attempt)
}