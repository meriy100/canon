package users

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	dbC "github.com/meriy100/canon/db"
	"github.com/meriy100/canon/models"
	"net/http"
	"strconv"
)

func Index(c echo.Context) error {
	db := dbC.GormConnect()
	defer db.Close()

	users := []models.User{}
	if err := db.Limit(10).Find(&users).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func Show(c echo.Context) error {
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
	if err := db.Create(&u).Error; err != nil {
		fmt.Println(err)
		return err
		//errorResponse := ErrorResponse{"Internal Server Error"}
		//return c.JSON(http.StatusInternalServerError, errorResponse)
	}
	return c.JSON(http.StatusOK, u)
}


