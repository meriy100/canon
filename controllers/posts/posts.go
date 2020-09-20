package posts

import (
	"github.com/labstack/echo"
	"github.com/meriy100/canon/application"
	"github.com/meriy100/canon/models"
	"net/http"
)

type Meta struct {
	Page uint `json:"page"`
	Per uint `json:"per"`
}

type Success struct {
	Data []models.Post `json:"data"`
	Meta Meta `json:"meta"`
}

func Index(c *application.Context) error {
	if _, err := c.AuthorizedUser(); err != nil {
		return echo.NewHTTPError(http.StatusForbidden)
	}

	posts := []models.Post{}
	if err := c.DB.Limit(10).Find(&posts).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, Success{ Data: posts, Meta: Meta{ Page: 1, Per: 10 } })
}
