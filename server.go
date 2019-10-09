package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"strconv"
	"github.com/meriy100/canon/db"
)

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "★This users id is " + id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

func save(c echo.Context) error {
	u := new(db.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func port() int {
	fmt.Println(os.Getenv("PORT"))
	if (os.Getenv("PORT") != "") {
		port, _ := strconv.Atoi(os.Getenv("PORT"))
		return port
	}
	return 1323
}

func main() {
	db := db.GormConnect()
	port := port()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!★★★★★")
	})
	e.GET("/users/:id", getUser)
	e.POST("/users", save)
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))

	defer db.Close()
}
