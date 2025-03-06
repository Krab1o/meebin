package auth

import (
	"github.com/Krab1o/meebin/internal/service"
	"github.com/go-playground/validator"
)

type handler struct {
	validate    *validator.Validate
	authService service.AuthService
}

func NewHandler(val *validator.Validate, as service.AuthService) *handler {
	return &handler{
		authService: as,
		validate:    val,
	}
}
