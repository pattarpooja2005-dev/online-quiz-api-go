package handlers

import (
	"time"
	"github.com/gin-gonic/gin"
	"quiz-api/config"
	"quiz-api/models"
)

func SubmitAnswers(c *gin.Context) {

	var answers []models.Answer

	if err := c.ShouldBindJSON(&answers); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	score := 0

	for _, ans := range answers {

		var question models.Question
		config.DB.First(&question, ans.QuestionID)

		if ans.Selected == question.Correct {
			score++
		}

		config.DB.Create(&ans)
	}

	attemptID := answers[0].AttemptID

	var attempt models.Attempt
	config.DB.First(&attempt, attemptID)

	attempt.Score = score
	attempt.EndTime = time.Now()
	config.DB.Save(&attempt)

	c.JSON(200, gin.H{"score": score})
}