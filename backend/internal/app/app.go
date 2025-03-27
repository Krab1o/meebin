package app

import (
	"context"
	"log"

	"github.com/Krab1o/meebin/internal/app/closer"
	"github.com/gin-gonic/gin"
)

const (
	envPath = ".env"
)

type App struct {
	serviceProvider *serviceProvider
	ginServer       *gin.Engine
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGinServer()
}

func (a *App) runGinServer() error {
	err := a.ginServer.Run(a.serviceProvider.HTTPConfig().Port())
	if err != nil {
		return err
	}
	return nil
}
