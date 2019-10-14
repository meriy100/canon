package sessionController

import (
	"github.com/labstack/echo"
	"github.com/meriy100/canon/application"
	"github.com/meriy100/canon/models"
	"net/http"
	"strconv"
)

type SessionForm struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Create(c *application.Context) error {
	u := models.User{}
	sf := SessionForm{}
	if err := c.Bind(&sf); err != nil {
		return c.String(http.StatusUnprocessableEntity, "missed email or password")
	}
	c.DB.Where("email = ?", sf.Email).First(&u)
	if !u.PasswordMach(sf.Password) {
		return c.String(http.StatusUnprocessableEntity, "missed email or password")
	}

	c.GetSession().Values["currentUserId"] = strconv.Itoa(int(u.ID))

	if err := c.GetSession().Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.String(http.StatusOK, "Logined")
}


func Destroy(c *application.Context) error {
	if _, err :=  c.AuthorizedUser(); err != nil {
		return echo.NewHTTPError(http.StatusForbidden)
	}
	c.GetSession().Values["currentUserId"] = ""
	if err := c.GetSession().Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.String(http.StatusOK, "Logouted")
}