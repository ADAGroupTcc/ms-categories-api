package categories

import (
	"net/http"

	"github.com/ADAGroupTcc/ms-categories-api/exceptions"
	"github.com/ADAGroupTcc/ms-categories-api/internal/domain"
	"github.com/ADAGroupTcc/ms-categories-api/internal/helpers"
	"github.com/ADAGroupTcc/ms-categories-api/internal/services/categories"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	Create(c echo.Context) error
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

func (h *categoriesHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	request := new(domain.CategoryRequest)
	if err := c.Bind(request); err != nil {
		return exceptions.New(exceptions.ErrInvalidPayload, err)
	}
	res, err := h.categoriesService.Create(ctx, request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, res)
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
	ctx := c.Request().Context()
	params := helpers.QueryParams{}
	if err := helpers.BindQueryParams(c, &params); err != nil {
		return exceptions.New(exceptions.ErrInvalidPayload, err)
	}

	var res *domain.CategoriesResponse
	var err error
	if len(params.CategoryIDs) < 1 {
		res, err = h.categoriesService.List(ctx, params.Limit, params.Offset)
		if err != nil {
			return err
		}
	} else {
		res, err = h.categoriesService.ListByCategoryIds(ctx, params.CategoryIDs, params.Limit, params.Offset)
		if err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, res)
}
