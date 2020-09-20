package users

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/meriy100/canon/application"
	"github.com/meriy100/canon/configs"
	"github.com/meriy100/canon/models"
	"net/http"
	"strconv"
	"time"
)

func Index(c *application.Context) error {
	if _, err := c.AuthorizedUser(); err != nil {
		return echo.NewHTTPError(http.StatusForbidden)
	}

	users := []models.User{}
	if err := c.DB.Limit(10).Find(&users).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func Show(c *application.Context) error {
	if _, err := c.AuthorizedUser(); err != nil {
		return echo.NewHTTPError(http.StatusForbidden)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	user.ID = uint(id)
	if err := c.DB.Find(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		return err
	}

	return c.JSON(http.StatusOK, user)
}


func Create(c *application.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	ep, _ := models.UserPassHash(u.Password)
	u.EncryptedPassword = ep
	if err := c.DB.Create(&u).Error; err != nil {
		return err
	}

	claims := &application.JwtCustomClaims{
		int(u.ID),
		jwt.StandardClaims{ ExpiresAt: time.Now().Add(time.Hour * 72).Unix() },
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(configs.GetSecretKey())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
