package categories

import (
	"net/http"

	"github.com/ADAGroupTcc/ms-categories-api/internal/services/categories"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetCategoryById(c echo.Context) error
	List(c echo.Context) error
}

type categoriesHandler struct {
	categoriesService categories.Service
}

func New(categoriesService categories.Service) Handler {
	return &categoriesHandler{
		categoriesService,
	}
}

func (h *categoriesHandler) GetCategoryById(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Param("id")
	res, err := h.categoriesService.GetCategoryById(ctx, userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *categoriesHandler) List(c echo.Context) error {
	return nil
}
