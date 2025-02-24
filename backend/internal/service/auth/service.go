package auth

import "github.com/Krab1o/meebin/internal/service"

type authService struct {
}

func NewService() service.AuthService {
	return &authService{}
}
