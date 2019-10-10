package router
import (
	"github.com/labstack/echo"
	"github.com/meriy100/canon/controllers/users"
	"net/http"
)

func Assign(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!★★★★★")
	})
	e.GET("/users", users.Index)
	e.GET("/users/:id", users.Show)
	e.POST("/users", users.Create)
}
