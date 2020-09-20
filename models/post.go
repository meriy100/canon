package models

import "time"

type Post struct {
	Id uint `json:"id" gorm: "primary key"`
	Title string `json:"string" gorm:"type: varchar(255) not null"`
	Body string `json:"string" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`
}
