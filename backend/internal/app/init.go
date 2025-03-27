package app

import (
	"context"
	"fmt"

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
			return fmt.Errorf("failed to init application: %w", err)
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(envPath)
	if err != nil {
		return fmt.Errorf("failed to init config: %w", err)
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
		return fmt.Errorf("failed to init gin server: %w", err)
	}
	a.SetupRoutes(ctx)
	return nil
}
