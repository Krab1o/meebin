package auth

import (
	"github.com/Krab1o/meebin/internal/config"
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Krab1o/meebin/internal/service"
)

type authService struct {
	userRepo    repository.UserRepository
	sessionRepo repository.UserSessionRepository
	jwtConf     config.JWTConfig
}

func NewService(
	userRepository repository.UserRepository,
	sessionRepository repository.UserSessionRepository,
	jwtConfig config.JWTConfig,
) service.AuthService {
	return &authService{
		userRepo:    userRepository,
		sessionRepo: sessionRepository,
		jwtConf:     jwtConfig,
	}
}
