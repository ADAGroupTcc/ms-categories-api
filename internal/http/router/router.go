package router

import (
	"fmt"

	"github.com/ADAGroupTcc/ms-categories-api/config"
	"github.com/ADAGroupTcc/ms-categories-api/internal/http/middlewares"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetupRouter(dependencies *config.Dependencies) *echo.Echo {
	e := echo.New()

	e.GET("/health", dependencies.HealthHandler.Check)
	fmt.Println(primitive.NewObjectID())

	v1 := e.Group("/v1")
	v1.POST("/categories", dependencies.Handler.Create, middlewares.ErrorIntercepter())
	v1.GET("/categories/:id", dependencies.Handler.GetCategoryById, middlewares.ErrorIntercepter())
	v1.GET("/categories", dependencies.Handler.List, middlewares.ErrorIntercepter())

	return e
}
