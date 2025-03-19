package event

import (
	"github.com/Krab1o/meebin/internal/service"
)

type Handler struct {
	eventService service.EventService
}

func NewHandler(es service.EventService) *Handler {
	return &Handler{eventService: es}
}
