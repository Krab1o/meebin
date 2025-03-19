package auth

import (
	"github.com/Krab1o/meebin/internal/service"
)

type Handler struct {
	authService service.AuthService
}

func NewHandler(as service.AuthService) *Handler {
	return &Handler{
		authService: as,
	}
}
