package sessionController

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	dbC "github.com/meriy100/canon/db"
	"github.com/labstack/echo"
	"github.com/meriy100/canon/models"
	"net/http"
	"strconv"
)

type SessionForm struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Create(c echo.Context) error {
	db := dbC.GormConnect()
	defer db.Close()
	u := models.User{}
	sf := SessionForm{}
	if err := c.Bind(&sf); err != nil {
		return c.String(http.StatusUnprocessableEntity, "missed email or password")
	}
	db.Where("email = ?", sf.Email).First(&u)
	if u.PasswordMach(sf.Password) {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["currentUserId"] = strconv.Itoa(int(u.ID))
		sess.Save(c.Request(), c.Response())

		return c.String(http.StatusOK, "Logined")
	} else {
		return c.String(http.StatusUnprocessableEntity, "missed email or password")
	}
}


func Destroy(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["currentUserId"] = ""
	sess.Save(c.Request(), c.Response())
	return c.String(http.StatusOK, "Logouted")
}