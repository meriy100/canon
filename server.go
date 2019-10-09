package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"strconv"
)
type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

func save(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func port() int {
	if (len(os.Args) > 1) {
		port, _ := strconv.Atoi(os.Args[1])
		return port
	}
	return 1323
}

func main() {
	port := port()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.GET("/users/:id", getUser)
	e.POST("/users", save)
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
