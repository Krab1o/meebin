package user

import (
	"github.com/Krab1o/meebin/internal/service"
)

type handler struct {
	service service.UserService
}

func NewAPI(us service.UserService) *handler {
	return &handler{service: us}
}
