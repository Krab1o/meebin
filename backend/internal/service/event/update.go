package event

import (
	"context"

	converter "github.com/Krab1o/meebin/internal/converter/service/event"
	rmodel "github.com/Krab1o/meebin/internal/model/event/r_model"
	smodel "github.com/Krab1o/meebin/internal/model/event/s_model"
	"github.com/Krab1o/meebin/internal/service"
)

func doUpdate(event *smodel.Event) bool {
	return event.Data != nil && *event.Data != smodel.EventData{} ||
		event.Status != 0
}

// TODO: add caller_id == jwtTokenUserId check if forbidden or not
func (s *serv) Update(
	ctx context.Context,
	updaterId uint64,
	event *smodel.Event,
) (*smodel.Event, error) {
	startUpdate := doUpdate(event)
	if !startUpdate {
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
		if event.Data != nil {
			err = s.eventRepository.UpdateEventData(ctx, repoEvent.Id, repoEvent.Data)
			if err != nil {
				return service.ErrorDBToService(err)
			}
		}
		if event.Status != 0 {
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
