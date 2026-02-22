package models

type Quiz struct {
	ID    uint   `gorm:"primaryKey"`
	Title string
}