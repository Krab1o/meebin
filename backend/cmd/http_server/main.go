package main

import (
	"context"
	"log"

	"github.com/Krab1o/meebin/internal/api"
	apiAuth "github.com/Krab1o/meebin/internal/api/auth"
	"github.com/Krab1o/meebin/internal/config"
	"github.com/Krab1o/meebin/internal/config/env"
	repoSession "github.com/Krab1o/meebin/internal/repository/session"
	repoUser "github.com/Krab1o/meebin/internal/repository/user"
	servAuth "github.com/Krab1o/meebin/internal/service/auth"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	path = ".env"
)

// TODO: replace error messages with strings
func main() {
	s := gin.Default()

	err := config.Load(path)
	if err != nil {
		log.Fatal(err)
	}

	httpConfig, err := env.NewHTTPConfig()
	if err != nil {
		log.Fatal("Failed to read httpConfig")
	}
	postgresConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatal("Failed to read pgConfig")
	}
	jwtConfig, err := env.NewJWTConfig()
	if err != nil {
		log.Println(err)
		log.Fatal("Failed to read jwtConfig")
	}

	//TODO: remove DB init to another function
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, postgresConfig.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("DB is not reachable, %v", err)
	}

	userRepository := repoUser.NewRepository(pool)
	sessionRepository := repoSession.NewRepository(pool)

	authService := servAuth.NewService(
		userRepository,
		sessionRepository,
		jwtConfig,
	)
	authHandler := apiAuth.NewHandler(authService)

	apiGroup := s.Group("/api")
	{
		authGroup := apiGroup.Group("/auth")
		{
			authGroup.POST("/register", api.MakeHandler(authHandler.Register)) // Регистрация
			authGroup.POST("/login", api.MakeHandler(authHandler.Login))       // Логин
			authGroup.POST("/refresh", api.MakeHandler(authHandler.Refresh))   // Обновление токена
			authGroup.POST("/logout", api.MakeHandler(authHandler.Logout))     // Выход
			authGroup.GET(
				"/profile",
				api.MakeHandler(authHandler.Profile),
			) // Данные пользователя
		}
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
	if err := s.Run(httpConfig.Port()); err != nil {
		log.Fatal(err)
	}
}
