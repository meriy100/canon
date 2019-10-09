package models

import "time"

type User struct {
	ID uint `json: "id" gorm:"primary_key"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null; unique"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`
}

