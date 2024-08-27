package categories

import (
	"github.com/ADAGroupTcc/ms-categories-api/internal/services/categories"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetCategoriesById(c echo.Context) error
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

func (h *categoriesHandler) GetCategoriesById(c echo.Context) error
func (h *categoriesHandler) List(c echo.Context) error
