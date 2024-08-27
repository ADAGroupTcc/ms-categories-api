package router

import (
	"github.com/ADAGroupTcc/ms-categories-api/config"
	"github.com/labstack/echo/v4"
)

func SetupRouter(dependencies *config.Dependencies) *echo.Echo {
	e := echo.New()

	e.GET("/health", dependencies.HealthHandler.Check)

	v1 := e.Group("/v1")
	v1.GET("/categories/:id", dependencies.Handler.GetCategoryById)
	v1.GET("/categories", dependencies.Handler.List)

	return e
}
