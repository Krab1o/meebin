package event

import (
	"context"

	"github.com/Krab1o/meebin/internal/service"
)

func (s *serv) Delete(ctx context.Context, deleterId uint64, eventId uint64) error {
	eventOwnerId, err := s.eventRepository.GetCallerIdById(ctx, eventId)
	if err != nil {
		return service.ErrorDBToService(err)
	}
	if deleterId != eventOwnerId {
		return service.NewForbiddenError(err)
	}
	err = s.eventRepository.DeleteEventById(ctx, eventId)
	if err != nil {
		return service.ErrorDBToService(err)
	}
	return nil
}
