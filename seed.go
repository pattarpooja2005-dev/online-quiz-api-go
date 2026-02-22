package main

import (
	"fmt"
	"quiz-api/config"
	"quiz-api/models"
)

func main() {

	// connect database
	config.ConnectDatabase()

	db := config.DB

	fmt.Println("Database connected")

	// ------------------------
	// CREATE QUIZ
	// ------------------------
	quiz := models.Quiz{
		Title:            "Math Quiz",
		TimeLimitMinutes: 10,
	}

	db.Create(&quiz)
	fmt.Println("Quiz created with ID:", quiz.ID)

	// ------------------------
	// CREATE QUESTIONS
	// ------------------------
	questions := []models.Question{
		{
			QuizID:  quiz.ID,
			Text:    "2 + 2 = ?",
			OptionA: "3",
			OptionB: "4",
			OptionC: "5",
			OptionD: "6",
			Correct: "B",
		},
		{
			QuizID:  quiz.ID,
			Text:    "5 × 3 = ?",
			OptionA: "10",
			OptionB: "12",
			OptionC: "15",
			OptionD: "20",
			Correct: "C",
		},
		{
			QuizID:  quiz.ID,
			Text:    "10 ÷ 2 = ?",
			OptionA: "2",
			OptionB: "3",
			OptionC: "4",
			OptionD: "5",
			Correct: "D",
		},
	}

	for _, q := range questions {
		db.Create(&q)
	}

	fmt.Println("Questions inserted successfully")
	fmt.Println("SEED COMPLETE")
}