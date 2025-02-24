package event

import "github.com/Krab1o/meebin/internal/service"

type eventService struct {}

func NewEventService() service.EventService{
	return &eventService{}
}