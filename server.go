package main

import (
	"fmt"
	dbC "github.com/meriy100/canon/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/meriy100/canon/application"
	"github.com/meriy100/canon/db"
	"github.com/meriy100/canon/router"
	"os"
	"strconv"
)

func port() int {
	fmt.Println(os.Getenv("PORT"))
	if (os.Getenv("PORT") != "") {
		port, _ := strconv.Atoi(os.Getenv("PORT"))
		return port
	}
	return 1323
}

func runServer() {
	db := dbC.GormConnect()
	port := port()
	e := echo.New()
	db.LogMode(true)
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            return h(&application.Context{c, db })
        }
    })
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	  Format: "[${method}] ${uri} : ${status} for ${time_rfc3339_nano}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
	router.Assign(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
	defer db.Close()
}

func main() {
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "migrate": db.Migration()
		case "drop": db.DropTables()
		default: runServer()
		}
	} else {
		runServer()
	}

}
