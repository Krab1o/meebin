package auth

import (
	"github.com/Krab1o/meebin/internal/config"
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Krab1o/meebin/internal/service"
)

type authService struct {
	userRepo    repository.UserRepository
	dataRepo    repository.UserDataRepository
	statsRepo   repository.UserStatsRepository
	sessionRepo repository.UserSessionRepository
	jwtConf     config.JWTConfig
}

func NewService(
	userRepository repository.UserRepository,
	dataRepository repository.UserDataRepository,
	statsRepository repository.UserStatsRepository,
	sessionRepository repository.UserSessionRepository,
	jwtConfig config.JWTConfig,
) service.AuthService {
	return &authService{
		userRepo:    userRepository,
		dataRepo:    dataRepository,
		statsRepo:   statsRepository,
		sessionRepo: sessionRepository,
		jwtConf:     jwtConfig,
	}
}
