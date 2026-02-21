package handlers

import (
	"github.com/gin-gonic/gin"
	"quiz-api/config"
	"quiz-api/models"
)

func CreateQuestion(c *gin.Context) {

	var question models.Question

	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&question)
	c.JSON(200, question)
}