package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Text        string `gorm:"not null"`
	IsCompleted bool   `gorm:"default:false"`
	ListID      uint   `gorm:"not null"`
}
