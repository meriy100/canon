package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/meriy100/canon/db"
	"github.com/meriy100/canon/router"
	"net/http"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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
		AllowOrigins: []string{"http://localhost:8080", "https://labstack.net"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.GET("/sessionExample", func(c echo.Context) error {
	  sess, _ := session.Get("session", c)
	  sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	  }
	  sess.Values["foo"] = "bar"
	  sess.Save(c.Request(), c.Response())
	  return c.String(http.StatusOK,  sess.Values["foo"].(string))
	})
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
