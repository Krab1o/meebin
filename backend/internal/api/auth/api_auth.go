package auth

import "github.com/Krab1o/meebin/internal/service"

type handler struct {
	authService service.AuthService
}

func NewHandler(as service.AuthService) *handler {
	return &handler{
		authService: as,
	}
}
