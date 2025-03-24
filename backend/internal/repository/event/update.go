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
		builder = builder.Set(rep.EventStatusColumn, event.Status)
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
func (r *repo) UpdateEventData(
	ctx context.Context,
	eventId uint64,
	eventData *rmodel.EventData,
) error {
	builder := sq.Update(rep.EventTableName).
		PlaceholderFormat(sq.Dollar)
	if eventData.Latitude != 0.0 {
		builder = builder.Set(rep.EventDataCallerIdColumn, eventData.Latitude)
	}
	if eventData.Longtitude != 0.0 {
		builder = builder.Set(rep.EventDataUtilizatorIdColumn, eventData.Longtitude)
	}
	if eventData.Title != "" {
		builder = builder.Set(rep.EventDataTitleColumn, eventData.Title)
	}
	if eventData.Description != "" {
		builder = builder.Set(rep.EventDataDescriptionColumn, eventData.Description)
	}
	if eventData.CallerId != 0 {
		builder = builder.Set(rep.EventDataCallerIdColumn, eventData.CallerId)
	}
	if eventData.UtilizatorId != 0 {
		builder = builder.Set(rep.EventDataUtilizatorIdColumn, eventData.UtilizatorId)
	}
	if time.Time.IsZero(eventData.TimeCalled) {
		builder = builder.Set(rep.EventDataTimeCalledColumn, eventData.TimeCalled)
	}
	if time.Time.IsZero(eventData.TimeUtilized) {
		builder = builder.Set(rep.EventDataTimeCleanedColumn, eventData.TimeUtilized)
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
