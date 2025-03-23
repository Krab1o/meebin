package app

import (
	"context"

	"github.com/Krab1o/meebin/docs"
	"github.com/Krab1o/meebin/internal/api"
	"github.com/Krab1o/meebin/internal/middleware"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Meebin
//	@version		0.1
//	@description	API for Meebin application
//	@contact.name	Krab1o
//	@contact.url	https://t.me/krab1o

//	@securityDefinitions.apikey	jwtToken
//	@in							header
//	@name						Authorization

//	@host		localhost:8080
//	@BasePath	/api
//	@accept		json
//	@produce	json
//	@schemes	http

//	@tag.name			Auth
//	@tag.description	Everything linked with authorization and JWT-token control

//	@tag.name			User
//	@tag.description	User control API

// @tag.name			Event
// @tag.description	Event control API
func (a *App) SetupRoutes(ctx context.Context) {
	docs.SwaggerInfo.BasePath = "/api"
	//TODO: add admin role to endpoints
	docsGroup := a.ginServer.Group("/docs")
	{
		docsGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	apiGroup := a.ginServer.Group("/api")
	{
		authGroup := apiGroup.Group("/auth")
		{
			// Регистрация
			authGroup.POST(
				"/register",
				api.MakeHandler(a.serviceProvider.AuthHandler(ctx).Register),
			)
			// Логин
			authGroup.POST(
				"/login",
				api.MakeHandler(a.serviceProvider.AuthHandler(ctx).Login),
			)
			// Обновление токена
			authGroup.POST(
				"/refresh",
				api.MakeHandler(a.serviceProvider.AuthHandler(ctx).Refresh),
			)
			// Логаут
			authGroup.POST(
				"/logout",
				api.MakeHandler(middleware.JWTMiddleware(a.serviceProvider.JWTConfig().Secret())),
				api.MakeHandler(a.serviceProvider.AuthHandler(ctx).Logout),
			)
		}
		userGroup := apiGroup.Group("/users")
		{
			// Получить список пользователей
			userGroup.GET(
				"",
				api.MakeHandler(middleware.JWTMiddleware(a.serviceProvider.JWTConfig().Secret())),
				api.MakeHandler(a.serviceProvider.UserHandler(ctx).ListUser),
			)
			// Получить пользователя по ID
			userGroup.GET(
				"/:id",
				api.MakeHandler(middleware.JWTMiddleware(a.serviceProvider.JWTConfig().Secret())),
				api.MakeHandler(a.serviceProvider.UserHandler(ctx).GetUser),
			)
			// Обновить данные пользователя
			userGroup.PATCH(
				"/:id",
				api.MakeHandler(middleware.JWTMiddleware(a.serviceProvider.JWTConfig().Secret())),
				api.MakeHandler(a.serviceProvider.UserHandler(ctx).UpdateUser),
			)
			// Удалить пользователя
			userGroup.DELETE(
				"/:id",
				api.MakeHandler(middleware.JWTMiddleware(a.serviceProvider.JWTConfig().Secret())),
				api.MakeHandler(a.serviceProvider.UserHandler(ctx).DeleteUser),
			)
		}
		events := apiGroup.Group("/events")
		{
			// Получить все события
			events.GET("", nil)
			// events.GET("", eventsHandler.ListEvents)
			// events.GET("/:id", eventsHandler.GetEvent)
			// events.POST("", eventsHandler.CreateEvent)
			// events.PATCH("/:id", eventsHandler.UpdateEvent)
			// events.DELETE("/:id", eventsHandler.DeleteEvent)
		}
	}
}
