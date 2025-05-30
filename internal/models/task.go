package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Completed   bool   `gorm:"not null" json:"completed"`
}
