package app

import (
	"context"
	"log"

	"github.com/Krab1o/meebin/internal/app/server"
	"github.com/Krab1o/meebin/internal/config"
	"github.com/gin-gonic/gin"
)

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGinServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(envPath)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGinServer(ctx context.Context) error {
	a.ginServer = gin.Default()
	err := server.ValidatorInit()
	if err != nil {
		log.Fatal(err)
	}
	a.SetupRoutes(ctx)
	return nil
}
