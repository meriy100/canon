package application

import (
	"github.com/jinzhu/gorm"
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/meriy100/canon/models"
	"strconv"
)

type Context struct {
	echo.Context
	Session *sessions.Session
	DB *gorm.DB
}


func (c Context) Foo() {
	fmt.Println("foo")
}


func(c *Context) GetSession() *sessions.Session {
	if c.Session != nil {
		return c.Session
	}
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	c.Session = sess

	return sess
}

func (c Context) AuthorizedUser() (models.User, error) {
	u := models.User{}
	value := c.GetSession().Values["currentUserId"]
	if value == nil {
		return u, errors.New("Session is empty")
	}
	userId, err := strconv.Atoi(value.(string))
	if err != nil {
		return u, errors.New("Session is empty")
	}
	u.ID = uint(userId)
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




