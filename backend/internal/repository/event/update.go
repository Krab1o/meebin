package event

import (
	"context"
	"time"

	rmodel "github.com/Krab1o/meebin/internal/model/event/r_model"
	rep "github.com/Krab1o/meebin/internal/repository"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) UpdateEvent(
	ctx context.Context,
	eventId uint64,
	event *rmodel.Event,
) error {
	builder := sq.Update(rep.EventTableName).
		PlaceholderFormat(sq.Dollar)
	if event.Status != 0 {
		builder = builder.Set(rep.EventColumnStatus, event.Status)
	}
	query, args, err := builder.ToSql()
	if err != nil {
		return rep.NewInternalError(err)
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return rep.NewInternalError(err)
	}

	return nil
}

// TODO: think how to refactor
// TODO: stackoverflow question
func (r *repo) UpdateEventData(
	ctx context.Context,
	eventId uint64,
	eventData *rmodel.EventData,
) error {
	builder := sq.Update(rep.EventTableName).
		PlaceholderFormat(sq.Dollar)
	if eventData.Latitude != 0.0 {
		builder = builder.Set(rep.EventDataColumnCallerId, eventData.Latitude)
	}
	if eventData.Longtitude != 0.0 {
		builder = builder.Set(rep.EventDataColumnUtilizatorId, eventData.Longtitude)
	}
	if eventData.Title != "" {
		builder = builder.Set(rep.EventDataColumnTitle, eventData.Title)
	}
	if eventData.Description != "" {
		builder = builder.Set(rep.EventDataColumnDescription, eventData.Description)
	}
	if eventData.CallerId != 0 {
		builder = builder.Set(rep.EventDataColumnCallerId, eventData.CallerId)
	}
	if eventData.UtilizatorId != 0 {
		builder = builder.Set(rep.EventDataColumnUtilizatorId, eventData.UtilizatorId)
	}
	if time.Time.IsZero(eventData.TimeCalled) {
		builder = builder.Set(rep.EventDataColumnTimeCalled, eventData.TimeCalled)
	}
	if time.Time.IsZero(eventData.TimeUtilized) {
		builder = builder.Set(rep.EventDataColumnTimeUtilized, eventData.TimeUtilized)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return rep.NewInternalError(err)
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return rep.NewInternalError(err)
	}
	return nil
}
