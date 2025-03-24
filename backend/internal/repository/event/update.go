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
	if event.CallerId != 0 {
		builder = builder.Set(rep.EventCallerIdColumn, event.CallerId)
	}
	if event.UtilizatorId != 0 {
		builder = builder.Set(rep.EventUtilizatorIdColumn, event.UtilizatorId)
	}
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

func (r *repo) UpdateEventData(
	ctx context.Context,
	eventId uint64,
	eventData *rmodel.EventData,
) error {
	builder := sq.Update(rep.EventTableName).
		PlaceholderFormat(sq.Dollar)
	if eventData.Latitude != 0.0 {
		builder = builder.Set(rep.EventCallerIdColumn, eventData.Latitude)
	}
	if eventData.Longtitude != 0.0 {
		builder = builder.Set(rep.EventUtilizatorIdColumn, eventData.Longtitude)
	}
	if eventData.Title != "" {
		builder = builder.Set(rep.EventDataTitleColumn, eventData.Title)
	}
	if eventData.Description != "" {
		builder = builder.Set(rep.EventDataDescriptionColumn, eventData.Description)
	}
	if time.Time.IsZero(eventData.TimeCalled) {
		builder = builder.Set(rep.EventDataTimeCalledColumn, eventData.TimeCalled)
	}
	if time.Time.IsZero(eventData.TimeCleaned) {
		builder = builder.Set(rep.EventDataTimeCleanedColumn, eventData.TimeCleaned)
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
