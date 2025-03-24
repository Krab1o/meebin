package event

import (
	"context"

	converter "github.com/Krab1o/meebin/internal/converter/service/event"
	smodel "github.com/Krab1o/meebin/internal/model/event/s_model"
	"github.com/Krab1o/meebin/internal/service"
)

func (s *serv) ListEvent(ctx context.Context) ([]smodel.Event, error) {
	repoEvents, err := s.eventRepository.List(ctx)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}
	events := make([]smodel.Event, len(repoEvents))
	for i, repoEvent := range repoEvents {
		events[i] = *converter.EventRepoToService(&repoEvent)
	}
	return events, nil
}
