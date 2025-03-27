package event

import (
	"github.com/Krab1o/meebin/internal/client/db"
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Krab1o/meebin/internal/service"
)

type serv struct {
	eventRepository repository.EventRepository
	txManager       db.TxManager
}

func NewService(
	eventRepository repository.EventRepository,
	txManager db.TxManager,
) service.EventService {
	return &serv{
		eventRepository: eventRepository,
		txManager:       txManager,
	}
}
