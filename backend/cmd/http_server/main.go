package main

import (
	"context"
	"log"

	"github.com/Krab1o/meebin/internal/api"
	apiAuth "github.com/Krab1o/meebin/internal/api/auth"
	apiUser "github.com/Krab1o/meebin/internal/api/user"
	"github.com/Krab1o/meebin/internal/config"
	"github.com/Krab1o/meebin/internal/config/env"
	"github.com/Krab1o/meebin/internal/middleware"
	repoRole "github.com/Krab1o/meebin/internal/repository/role"
	repoSession "github.com/Krab1o/meebin/internal/repository/session"
	repoUser "github.com/Krab1o/meebin/internal/repository/user"
	servAuth "github.com/Krab1o/meebin/internal/service/auth"
	servUser "github.com/Krab1o/meebin/internal/service/user"
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
		log.Fatal("Failed to setup httpConfig")
	}
	postgresConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatal("Failed to setup pgConfig")
	}
	jwtConfig, err := env.NewJWTConfig()
	if err != nil {
		log.Fatal("Failed to setup jwtConfig")
	}

	//TODO: remove DB init to DI container
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

	//TODO: move validator init to DI container
	//TODO: inverse dependencies somehow
	err = validatorInit()
	if err != nil {
		// log.Fatal("Failed to setup validator")
	}

	userRepository := repoUser.NewRepository(pool)
	sessionRepository := repoSession.NewRepository(pool)
	roleRepository := repoRole.NewRepository(pool)
	authService := servAuth.NewService(
		userRepository,
		sessionRepository,
		roleRepository,
		jwtConfig,
	)
	userService := servUser.NewService(userRepository)

	authHandler := apiAuth.NewHandler(authService)
	userHandler := apiUser.NewHandler(userService)

	//TODO: add admin role to endpoints
	apiGroup := s.Group("/api")
	{
		authGroup := apiGroup.Group("/auth")
		{
			// Регистрация
			authGroup.POST(
				"/register",
				api.MakeHandler(authHandler.Register),
			)
			// Логин
			authGroup.POST(
				"/login",
				api.MakeHandler(authHandler.Login),
			)
			// Обновление токена
			authGroup.POST(
				"/refresh",
				api.MakeHandler(authHandler.Refresh),
			)
			// Логаут
			authGroup.POST(
				"/logout",
				api.MakeHandler(middleware.JWTMiddleware(jwtConfig.Secret())),
				api.MakeHandler(authHandler.Logout),
			)
		}
		userGroup := apiGroup.Group("/users")
		{
			// Получить список пользователей
			userGroup.GET(
				"",
				api.MakeHandler(middleware.JWTMiddleware(jwtConfig.Secret())),
				api.MakeHandler(userHandler.ListUser),
			)
			// Получить пользователя по ID
			userGroup.GET(
				"/:id",
				api.MakeHandler(middleware.JWTMiddleware(jwtConfig.Secret())),
				api.MakeHandler(userHandler.GetUser),
			)
			// Обновить данные пользователя
			userGroup.PATCH(
				"/:id",
				api.MakeHandler(middleware.JWTMiddleware(jwtConfig.Secret())),
				api.MakeHandler(userHandler.UpdateUser),
			)
			// Удалить пользователя
			userGroup.DELETE(
				"/:id",
				api.MakeHandler(middleware.JWTMiddleware(jwtConfig.Secret())),
				api.MakeHandler(userHandler.DeleteUser),
			)
		}
		events := apiGroup.Group("/events")
		{
			// Получить все события
			events.GET("", nil)
		}
	}

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
