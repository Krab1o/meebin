package event

import (
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Krab1o/meebin/internal/service"
)

type eventService struct {
	eventRepo repository.EventRepository
}

func NewService(eventRepository repository.EventRepository) service.EventService {
	return &eventService{
		eventRepo: eventRepository,
	}
}
