package event

import (
	"context"

	"github.com/Krab1o/meebin/internal/service"
)

func (s *serv) Delete(ctx context.Context, eventId uint64) error {
	err := s.eventRepository.DeleteById(ctx, eventId)
	if err != nil {
		return service.ErrorDBToService(err)
	}
	return nil
}
