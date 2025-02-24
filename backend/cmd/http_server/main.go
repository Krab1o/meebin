package main

import (
	"log"

	hAuth "github.com/Krab1o/meebin/internal/api/auth"
	sAuth "github.com/Krab1o/meebin/internal/service/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()
	SetupRoutes(s)

	api := g.Group("/api")
	authService := sAuth.NewService()
	authHandler := hAuth.NewHandler(authService)
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register) // Регистрация
		// auth.POST("/login", authHandler.Login)        // Логин
		// auth.POST("/refresh", authHandler.Refresh)    // Обновление токена
		// auth.POST("/logout", authHandler.Logout)      // Выход
		// auth.GET("/profile", authHandler.Profile)       // Данные пользователя
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
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
