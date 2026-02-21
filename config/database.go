package config

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"quiz-api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("quiz.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = database

	DB.AutoMigrate(
		&models.Quiz{},
		&models.Question{},
		&models.Attempt{},
		&models.Answer{},
	)

	fmt.Println("Database connected")
}