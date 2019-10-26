package sessionController

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/meriy100/canon/application"
	"github.com/meriy100/canon/models"
	"net/http"
	"time"
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
	claims := &application.JwtCustomClaims{
		int(u.ID),
		jwt.StandardClaims{ ExpiresAt: time.Now().Add(time.Hour * 72).Unix() },
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
