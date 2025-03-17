package user

import (
	"github.com/Krab1o/meebin/internal/service"
)

type handler struct {
	userService service.UserService
}

func NewHandler(us service.UserService) *handler {
	return &handler{
		userService: us,
	}
}
