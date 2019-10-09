package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/meriy100/canon/router"
	"os"
	"strconv"
	"github.com/meriy100/canon/db"
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
