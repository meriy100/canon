package application

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/meriy100/canon/models"
)

type JwtCustomClaims struct {
	IDHash int `json:"int"`
	jwt.StandardClaims
}


type Context struct {
	echo.Context
	DB *gorm.DB
}


func (c Context) Foo() {
	fmt.Println("foo")
}

func (c Context) AuthorizedUser() (models.User, error) {
	u := models.User{}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	IDHash := claims.IDHash

	u.ID = uint(IDHash)
	if err := c.DB.Find(&u).Error; err != nil {
		fmt.Println(err)
		return u, errors.New("User is not valid")
	}
	return u, nil
}

type callFunc func(c *Context) error

func CallHandler(h callFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        return h(c.(*Context))
    }
}




