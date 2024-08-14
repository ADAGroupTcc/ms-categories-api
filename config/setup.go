package config

import (
	"context"

	handler "github.com/ADAGroupTcc/ms-categories-api/internal/http/categories"
	"github.com/ADAGroupTcc/ms-categories-api/internal/http/health"
	repository "github.com/ADAGroupTcc/ms-categories-api/internal/repositories/categories"
	service "github.com/ADAGroupTcc/ms-categories-api/internal/services/categories"
	healthService "github.com/ADAGroupTcc/ms-categories-api/internal/services/health"
	"github.com/ADAGroupTcc/ms-categories-api/pkg/mongorm"
)

type Dependencies struct {
	Handler       handler.Handler
	HealthHandler health.Health
}

func NewDependencies(ctx context.Context, envs *Environments) *Dependencies {
	database, err := mongorm.Connect(envs.DBUri, envs.DBName)
	if err != nil {
		panic(err)
	}
	userRepository := repository.New(database)
	userService := service.New(userRepository)
	userHandler := handler.New(userService)

	healthService := healthService.New(database)
	healthHandler := health.New(healthService)
	return &Dependencies{
		userHandler,
		healthHandler,
	}
}
