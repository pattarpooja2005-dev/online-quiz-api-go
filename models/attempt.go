package models

import "time"

type Attempt struct {
	ID        uint `gorm:"primaryKey"`
	QuizID    uint
	UserName  string
	StartTime time.Time
}