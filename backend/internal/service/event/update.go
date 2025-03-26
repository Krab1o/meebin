package event

import (
	"context"

	converter "github.com/Krab1o/meebin/internal/converter/service/event"
	rmodel "github.com/Krab1o/meebin/internal/model/event/r_model"
	smodel "github.com/Krab1o/meebin/internal/model/event/s_model"
	"github.com/Krab1o/meebin/internal/service"
)

func doDataUpdate(event *smodel.Event) bool {
	return event.Data != nil && *event.Data != smodel.EventData{}
}

func doStatusUpdate(event *smodel.Event) bool {
	return event.Status != 0
}

func (s *serv) Update(
	ctx context.Context,
	updaterId uint64,
	event *smodel.Event,
) (*smodel.Event, error) {
	if updaterId != event.Data.CallerId {
		return nil, service.NewForbiddenError(nil)
	}

	dataUpdate := doDataUpdate(event)
	statusUpdate := doStatusUpdate(event)
	if !dataUpdate && !statusUpdate {
		return nil, service.NewNoUpdateError(nil)
	}

	eventOwnerId, err := s.eventRepository.GetCallerIdById(ctx, event.Id)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}
	if updaterId != eventOwnerId {
		return nil, service.NewForbiddenError(err)
	}

	var updatedRepoEvent *rmodel.Event
	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		repoEvent := converter.EventServiceToRepo(event)
		var err error
		if dataUpdate {
			err = s.eventRepository.UpdateEventData(ctx, repoEvent.Id, repoEvent.Data)
			if err != nil {
				return service.ErrorDBToService(err)
			}
		}
		if statusUpdate {
			err = s.eventRepository.UpdateEvent(ctx, repoEvent.Id, repoEvent)
			if err != nil {
				return service.ErrorDBToService(err)
			}
		}
		updatedRepoEvent, err = s.eventRepository.GetEventById(ctx, event.Id)
		if err != nil {
			return service.ErrorDBToService(err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	updatedEvent := converter.EventRepoToService(updatedRepoEvent)
	return updatedEvent, nil
}
