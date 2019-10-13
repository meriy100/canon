package users

import (
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	dbC "github.com/meriy100/canon/db"
	"github.com/meriy100/canon/models"
	"net/http"
	"strconv"
)

func Index(c echo.Context) error {
	if _, err := authorizedUser(c); err != nil {
		return echo.NewHTTPError(http.StatusForbidden)
	}
	db := dbC.GormConnect()
	defer db.Close()

	users := []models.User{}
	if err := db.Limit(10).Find(&users).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func Show(c echo.Context) error {
	if _, err := authorizedUser(c); err != nil {
		return echo.NewHTTPError(http.StatusForbidden)
	}
	db := dbC.GormConnect()
	defer db.Close()

	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	user.ID = uint(id)
	if err := db.Find(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		return err
	}

	return c.JSON(http.StatusOK, user)
}


func Create(c echo.Context) error {
	db := dbC.GormConnect()
	defer db.Close()

	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	ep, _ := models.UserPassHash(u.Password)
	u.EncryptedPassword = ep
	if err := db.Create(&u).Error; err != nil {
		fmt.Println(err)
		return err
		//errorResponse := ErrorResponse{"Internal Server Error"}
		//return c.JSON(http.StatusInternalServerError, errorResponse)
	}
	u.Password = ""
	u.PasswordConfirmation = ""
	return c.JSON(http.StatusOK, u)
}



func authorizedUser(c echo.Context) (models.User, error) {
	u := models.User{}
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	db := dbC.GormConnect()
	defer db.Close()
	userId, err := strconv.Atoi(sess.Values["currentUserId"].(string))
	if err != nil {
		return u, errors.New("Session is empty")
	}
	u.ID = uint(userId)
	if err := db.Find(&u).Error; err != nil {
		fmt.Println(err)
		return u, errors.New("User is not valid")
	}
	return u, nil
}
