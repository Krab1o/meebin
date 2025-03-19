package app

import (
	"context"
	"log"

	apiAuth "github.com/Krab1o/meebin/internal/api/auth"
	apiEvent "github.com/Krab1o/meebin/internal/api/event"
	apiUser "github.com/Krab1o/meebin/internal/api/user"
	"github.com/Krab1o/meebin/internal/app/closer"
	"github.com/Krab1o/meebin/internal/config"
	"github.com/Krab1o/meebin/internal/config/env"
	"github.com/Krab1o/meebin/internal/repository"
	repoEvent "github.com/Krab1o/meebin/internal/repository/event"
	repoRole "github.com/Krab1o/meebin/internal/repository/role"
	repoSession "github.com/Krab1o/meebin/internal/repository/session"
	repoUser "github.com/Krab1o/meebin/internal/repository/user"
	"github.com/Krab1o/meebin/internal/service"
	servAuth "github.com/Krab1o/meebin/internal/service/auth"
	servEvent "github.com/Krab1o/meebin/internal/service/event"
	servUser "github.com/Krab1o/meebin/internal/service/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	httpConfig config.HTTPConfig
	jwtConfig  config.JWTConfig

	pgPool *pgxpool.Pool

	userRepo    repository.UserRepository
	eventRepo   repository.EventRepository
	roleRepo    repository.RoleRepository
	sessionRepo repository.SessionRepository

	authService  service.AuthService
	userService  service.UserService
	eventService service.EventService

	authHandler  *apiAuth.Handler
	userHandler  *apiUser.Handler
	eventHandler *apiEvent.Handler
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatal("Failed to setup pgConfig")
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatal("Failed to setup httpConfig")
		}
		s.httpConfig = cfg
	}
	return s.httpConfig
}

func (s *serviceProvider) JWTConfig() config.JWTConfig {
	if s.jwtConfig == nil {
		cfg, err := env.NewJWTConfig()
		if err != nil {
			log.Fatal("Failed to setup jwtConfig")
		}
		s.jwtConfig = cfg
	}
	return s.jwtConfig
}

func (s *serviceProvider) PGPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.New(ctx, s.pgConfig.DSN())
		if err != nil {
			log.Fatalf("Failed to connect to DB: %v", err)
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})
		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("DB is not reachable, %v", err)
		}
		s.pgPool = pool
	}
	return s.pgPool
}

// 	sessionRepository := repoSession.NewRepository(pool)
// 	roleRepository := repoRole.NewRepository(pool)
// 	authService := servAuth.NewService(
// 		userRepository,
// 		sessionRepository,
// 		roleRepository,
// 		jwtConfig,
// 	)
// 	userService := servUser.NewService(userRepository)

// 	authHandler := apiAuth.NewHandler(authService)
// 	userHandler := apiUser.NewHandler(userService)

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepo == nil {
		s.userRepo = repoUser.NewRepository(s.pgPool)
	}
	return s.userRepo
}

func (s *serviceProvider) EventRepository(ctx context.Context) repository.EventRepository {
	if s.eventRepo == nil {
		s.eventRepo = repoEvent.NewRepository(s.pgPool)
	}
	return s.eventRepo
}

func (s *serviceProvider) SessionRepository(ctx context.Context) repository.SessionRepository {
	if s.sessionRepo == nil {
		s.sessionRepo = repoSession.NewRepository(s.pgPool)
	}
	return s.sessionRepo
}

func (s *serviceProvider) RoleRepository(ctx context.Context) repository.RoleRepository {
	if s.roleRepo == nil {
		s.roleRepo = repoRole.NewRepository(s.pgPool)
	}
	return s.roleRepo
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = servAuth.NewService(
			s.UserRepository(ctx),
			s.SessionRepository(ctx),
			s.RoleRepository(ctx),
			s.JWTConfig(),
		)
	}
	return s.authService
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = servUser.NewService(s.UserRepository(ctx))
	}
	return s.userService
}

func (s *serviceProvider) EventService(ctx context.Context) service.EventService {
	if s.eventService == nil {
		s.eventService = servEvent.NewService(s.EventRepository(ctx))
	}
	return s.eventService
}

func (s *serviceProvider) AuthHandler(ctx context.Context) *apiAuth.Handler {
	if s.authHandler == nil {
		s.authHandler = apiAuth.NewHandler(s.AuthService(ctx))
	}
	return s.authHandler
}

func (s *serviceProvider) UserHandler(ctx context.Context) *apiUser.Handler {
	if s.userHandler == nil {
		s.userHandler = apiUser.NewHandler(s.UserService(ctx))
	}
	return s.userHandler
}

func (s *serviceProvider) Eventhandler(ctx context.Context) *apiEvent.Handler {
	if s.eventHandler == nil {
		s.eventHandler = apiEvent.NewHandler(s.EventService(ctx))
	}
	return s.eventHandler
}
