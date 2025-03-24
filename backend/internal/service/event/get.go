package event

import (
	"context"

	converter "github.com/Krab1o/meebin/internal/converter/service/event"
	smodel "github.com/Krab1o/meebin/internal/model/event/s_model"
	"github.com/Krab1o/meebin/internal/service"
)

func (s *serv) GetEvent(ctx context.Context, eventId uint64) (*smodel.Event, error) {
	repoEvent, err := s.eventRepository.GetById(ctx, eventId)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}
	event := converter.EventRepoToService(repoEvent)
	return event, nil
}
