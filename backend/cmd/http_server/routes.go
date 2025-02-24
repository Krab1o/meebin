package main

import (
	"github.com/Krab1o/meebin/internal/api/auth"
	"github.com/gin-gonic/gin"
)

func addAPIEndpoints(api *gin.RouterGroup) {
	authHandler := auth.NewHandler()
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)  // Регистрация
		auth.POST("/login", authHandler.Login)        // Логин
		auth.POST("/refresh", authHandler.Refresh)    // Обновление токена
		auth.POST("/logout", authHandler.Logout)      // Выход
		auth.GET("/profile", authHandler.Profile)       // Данные пользователя
	}
	// users := api.Group("/users")
	// usersHandler := user.NewHandler()
	// users.GET("", usersHandler.ListUser)
	// users.GET("/:id", usersHandler.GetUser)
	// users.POST("", usersHandler.CreateUser)
	// users.PATCH("/:id", usersHandler.UpdateUser)
	// users.DELETE("/:id", usersHandler.DeleteUser)
	
	// events := api.Group("/events")
	// eventsHandler := event.NewHandler()
	// events.GET("", eventsHandler.ListEvents)
	// events.GET("/:id", eventsHandler.GetEvent)
	// events.POST("", eventsHandler.CreateEvent)
	// events.PATCH("/:id", eventsHandler.UpdateEvent)
	// events.DELETE("/:id", eventsHandler.DeleteEvent)
}

func SetupRoutes(g *gin.Engine) {
	api := g.Group("/api")
	addAPIEndpoints(api)
}