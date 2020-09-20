package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/meriy100/canon/application"
	"github.com/meriy100/canon/configs"
	"github.com/meriy100/canon/controllers/posts"
	"github.com/meriy100/canon/controllers/sessionController"
	"github.com/meriy100/canon/controllers/users"
)

func Assign(e *echo.Echo) {
	config := middleware.JWTConfig{
		Claims:     &application.JwtCustomClaims{},
		SigningKey: configs.GetSecretKey(),
	}
	g := e.Group("/")
	g.Use(middleware.JWTWithConfig(config))
	g.GET("users", application.CallHandler(users.Index))
	g.GET("users/:id", application.CallHandler(users.Show))

	g.GET("posts", application.CallHandler(posts.Index))

	//e.POST("/sign_up", application.CallHandler(users.Create))
	e.POST("/session", application.CallHandler(sessionController.Create))
}
