package auth

import (
	"github.com/Krab1o/meebin/internal/config"
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Krab1o/meebin/internal/service"
)

type serv struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	roleRepo    repository.RoleRepository
	jwtConf     config.JWTConfig
}

func NewService(
	userRepository repository.UserRepository,
	sessionRepository repository.SessionRepository,
	roleRepository repository.RoleRepository,
	jwtConfig config.JWTConfig,
) service.AuthService {
	return &serv{
		userRepo:    userRepository,
		sessionRepo: sessionRepository,
		roleRepo:    roleRepository,
		jwtConf:     jwtConfig,
	}
}
