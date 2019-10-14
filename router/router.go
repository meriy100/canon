package router

import (
	"github.com/labstack/echo"
	"github.com/meriy100/canon/application"
	"github.com/meriy100/canon/controllers/sessionController"
	"github.com/meriy100/canon/controllers/users"
)

func Assign(e *echo.Echo) {
	e.GET("/users", application.CallHandler(users.Index))
	e.GET("/users/:id", application.CallHandler(users.Show))
	e.POST("/users", application.CallHandler(users.Create))

	e.POST("/session", application.CallHandler(sessionController.Create))
	e.DELETE("/session", application.CallHandler(sessionController.Destroy))
}
