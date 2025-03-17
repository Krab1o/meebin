package event

import (
	"github.com/Krab1o/meebin/internal/service"
)

type handler struct {
	eventService service.EventService
}

func NewHandler(es service.EventService) *handler {
	return &handler{eventService: es}
}