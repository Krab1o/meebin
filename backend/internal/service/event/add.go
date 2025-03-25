package event

import (
	"context"

	converter "github.com/Krab1o/meebin/internal/converter/service/event"
	"github.com/Krab1o/meebin/internal/model"
	smodel "github.com/Krab1o/meebin/internal/model/event/s_model"
	"github.com/Krab1o/meebin/internal/service"
)

func (s *serv) Create(ctx context.Context, event *smodel.Event) (uint64, error) {
	event.Status = model.StatusOnModeration
	repoEvent := converter.EventServiceToRepo(event)

	var eventId uint64
	var err error
	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		eventId, err = s.eventRepository.AddEvent(ctx, repoEvent)
		if err != nil {
			return service.ErrorDBToService(err)
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return eventId, err
}
