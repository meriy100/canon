package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID uint `json: "id" gorm:"primary_key"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null; unique"`
	//Password string `json:"email" gorm:"not null; unique"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`
}


func (user User) Validate(db *gorm.DB) {
	if len(user.Email) == 0 {
		db.AddError(errors.New("email is blank"))
	}
	if len(user.Name) == 0 {
		db.AddError(errors.New("name is blank"))
	}
}
