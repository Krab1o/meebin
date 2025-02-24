package user

import "github.com/Krab1o/meebin/internal/service"

type userService struct{}

func NewService() service.UserService {
	return &userService{}
}
