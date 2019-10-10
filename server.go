package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/meriy100/canon/db"
	"github.com/meriy100/canon/router"

	"github.com/labstack/echo/middleware"
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
	port := port()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	router.Assign(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
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
