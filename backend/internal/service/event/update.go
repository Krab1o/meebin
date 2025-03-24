package event

import (
	"context"

	converter "github.com/Krab1o/meebin/internal/converter/service/event"
	smodel "github.com/Krab1o/meebin/internal/model/event/s_model"
	"github.com/Krab1o/meebin/internal/service"
)

func doUpdate(event *smodel.Event) bool {
	return event.Data != nil && *event.Data != smodel.EventData{} ||
		event.Status != 0
}

func (s *serv) Update(ctx context.Context, eventId uint64, event *smodel.Event) error {
	startUpdate := doUpdate(event)
	if !startUpdate {
		return service.NewNoUpdateError(nil)
	}

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		repoEvent := converter.EventServiceToRepo(event)
		var err error
		if event.Data != nil {
			err = s.eventRepository.UpdateEventData(ctx, eventId, repoEvent.Data)
			if err != nil {
				return service.ErrorDBToService(err)
			}
		}
		if event.Status != 0 {
			err = s.eventRepository.UpdateEvent(ctx, eventId, repoEvent)
			if err != nil {
				return service.ErrorDBToService(err)
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
