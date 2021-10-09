package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	Status   bool   `json:"status" gorm:"default:true"`
}
