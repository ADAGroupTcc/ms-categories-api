package main

import (
	"context"

	"github.com/ADAGroupTcc/ms-categories-api/config"
	"github.com/ADAGroupTcc/ms-categories-api/internal/http/router"
)

func main() {
	ctx := context.TODO()
	envs, err := config.LoadEnvVars()
	if err != nil {
		panic(err)
	}
	dependencies := config.NewDependencies(ctx, envs)
	e := router.SetupRouter(dependencies)
	err = e.Start(":" + envs.ApiPort)
	if err != nil {
		panic(err)
	}
}
