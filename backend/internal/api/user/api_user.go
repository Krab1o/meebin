package user

import (
	"github.com/Krab1o/meebin/internal/service"
)

type Handler struct {
	userService service.UserService
}

func NewHandler(us service.UserService) *Handler {
	return &Handler{
		userService: us,
	}
}
