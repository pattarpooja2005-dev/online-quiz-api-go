package models

type Question struct {
	ID        uint   `gorm:"primaryKey"`
	QuizID    uint
	Text      string
	OptionA   string
	OptionB   string
	OptionC   string
	OptionD   string
	Correct   string
}
