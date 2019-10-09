package users

import (
	"github.com/labstack/echo"
	dbC "github.com/meriy100/canon/db"
	"net/http"
	"strconv"
)

func Index(c echo.Context) error {
	db := dbC.GormConnect()
	defer db.Close()

	users := []dbC.User{}
	db.Limit(10).Find(&users)
	return c.JSON(http.StatusCreated, users)
}

func Show(c echo.Context) error {
	db := dbC.GormConnect()
	defer db.Close()

	id, _ := strconv.Atoi(c.Param("id"))
	user := dbC.User{}
	user.ID = uint(id)
	db.Find(&user)

	return c.JSON(http.StatusOK, user)
}


func Create(c echo.Context) error {
	db := dbC.GormConnect()
	defer db.Close()

	u := new(dbC.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	db.Create(&u)
	return c.JSON(http.StatusOK, u)
}


