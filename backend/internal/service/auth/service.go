package auth

import (
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Krab1o/meebin/internal/service"
)

type authService struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
}

func NewService(
	userRepository repository.UserRepository,
	sessionRepository repository.SessionRepository,
) service.AuthService {
	return &authService{
		userRepo:    userRepository,
		sessionRepo: sessionRepository,
	}
}
