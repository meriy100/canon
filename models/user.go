package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
	"golang.org/x/crypto/bcrypt"
)

type UserForm struct {
	User
	PasswordPair
}

type User struct {
	ID uint `json: "id" gorm:"primary_key"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null; unique"`
	EncryptedPassword string `json:"-" gorm:"not null;"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`
	Password string `json:"password,omitempty" gorm:"-"'`
	PasswordConfirmation string `json:"passwordConfirmation,omitempty" gorm:"-"`
}


func UserValidate(user User, db *gorm.DB) {
	if len(user.Email) == 0 {
		db.AddError(errors.New("email is blank"))
	}
	if len(user.Name) == 0 {
		db.AddError(errors.New("name is blank"))
	}
	if len(user.EncryptedPassword) == 0 {
		db.AddError(errors.New("password is blank"))
	}
	if len(user.Password) < 8 {
		db.AddError(errors.New("password is over 8 characters"))
	}
	if user.Password != user.PasswordConfirmation {
		db.AddError(errors.New("password is not same password confirmation"))
	}
}


type PasswordPair struct {
	Password string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

func PasswordPairValidate(pp PasswordPair, db *gorm.DB) {
}

func (user User) Validate(db *gorm.DB) {
	UserValidate(user, db)
}

func (pp PasswordPair) Validate(db *gorm.DB) {
	PasswordPairValidate(pp, db)
}

func (user User) PasswordMach(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(pw)) == nil
}

func UserPassHash(pass string) (string, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}