package auth

import (
	"github.com/Krab1o/meebin/internal/client/db"
	"github.com/Krab1o/meebin/internal/config"
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Krab1o/meebin/internal/service"
)

type serv struct {
	userRepository    repository.UserRepository
	sessionRepository repository.SessionRepository
	roleRepository    repository.RoleRepository
	jwtConfig         config.JWTConfig
	txManager         db.TxManager
}

func NewService(
	userRepository repository.UserRepository,
	sessionRepository repository.SessionRepository,
	roleRepository repository.RoleRepository,
	jwtConfig config.JWTConfig,
	txManager db.TxManager,
) service.AuthService {
	return &serv{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
		roleRepository:    roleRepository,
		jwtConfig:         jwtConfig,
		txManager:         txManager,
	}
}
