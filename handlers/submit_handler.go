package handlers

import (
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

	config.DB.Create(&answers)

	score := calculateScore(answers)

	c.JSON(200, gin.H{
		"score": score,
	})
}

func calculateScore(answers []models.Answer) int {

	score := 0

	for _, ans := range answers {

		var q models.Question
		config.DB.First(&q, ans.QuestionID)

		if ans.Selected == q.Correct {
			score++
		}
	}

	return score
}