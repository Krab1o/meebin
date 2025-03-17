package user

import (
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Krab1o/meebin/internal/service"
)

type serv struct {
	userRepo repository.UserRepository
}

func NewService(userRepository repository.UserRepository) service.UserService {
	return &serv{
		userRepo: userRepository,
	}
}
