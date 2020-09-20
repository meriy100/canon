package posts

import (
	"github.com/labstack/echo"
	"github.com/meriy100/canon/application"
	"github.com/meriy100/canon/models"
	"net/http"
)

type Meta struct {
	models.Pagination
}

type Success struct {
	Data []models.Post `json:"data"`
	Meta Meta `json:"meta"`
}

func Index(c *application.Context) error {
	if _, err := c.AuthorizedUser(); err != nil {
		return echo.NewHTTPError(http.StatusForbidden)
	}

	pagination := new(models.Pagination)

	if err := c.Bind(pagination); err != nil {
		return err
	}

	var count uint
	posts := []models.Post{}
	if err := c.DB.Model(&models.Post{}).Count(&count).Error; err != nil {
		return err
	}

	if err := c.DB.Limit(pagination.Per).Offset(models.ToOffset(pagination)).Find(&posts).Error; err != nil {
		return err
	}
	return c.JSON(
		http.StatusOK, Success{
			Data: posts,
			Meta: Meta{
				models.Pagination{
					Page: 1,
					Per: 10,
					Count: count } } } )
}
