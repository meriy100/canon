package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func GormConnect() *gorm.DB {
	DBMS     := "postgres"
	USER     := "canon_user"
	PASS     := "password"
	HOST     := "localhost"
	//PORT := ""
	DBNAME   := "canon"

	//CONNECT := "host=" + HOST + " port=" + PORT + " user=" + USER + "password=" + PASS + " dbname=" + DBNAME
	CONNECT := "host=" + HOST + " user=" + USER + " password=" + PASS + " dbname=" + DBNAME + " sslmode=disable"
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}


type User struct {
	Id int `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func migration() {
	db := GormConnect()
	db.AutoMigrate(&User{})
}

