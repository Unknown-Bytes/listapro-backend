package models

import "gorm.io/gorm"

type TaskList struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Tasks []Task `gorm:"foreignkey:ListID"`
}
