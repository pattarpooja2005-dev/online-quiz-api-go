package models

type Answer struct {
	ID         uint   `gorm:"primaryKey"`
	AttemptID  uint
	QuestionID uint
	Selected   string
}