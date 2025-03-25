package event

import (
	"context"
	"errors"
	"fmt"

	rmodel "github.com/Krab1o/meebin/internal/model/event/r_model"
	rep "github.com/Krab1o/meebin/internal/repository"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) GetEventById(ctx context.Context, eventId uint64) (*rmodel.Event, error) {
	query, args, err := sq.Select(
		rep.Col(rep.EventTableName, rep.EventIdColumn),

		rep.Col(rep.EventDataTableName, rep.EventDataCallerIdColumn),
		rep.Col(rep.EventDataTableName, rep.EventDataUtilizatorIdColumn),
		rep.Col(rep.EventDataTableName, rep.EventDataLatitudeColumn),
		rep.Col(rep.EventDataTableName, rep.EventDataLongtitudeColumn),
		rep.Col(rep.EventDataTableName, rep.EventDataTitleColumn),
		rep.Col(rep.EventDataTableName, rep.EventDataDescriptionColumn),
		rep.Col(rep.EventDataTableName, rep.EventDataTimeCalledColumn),
		rep.Col(rep.EventDataTableName, rep.EventDataTimeUtilizedColumn),

		rep.Col(rep.EventStatusTableName, rep.EventStatusTitleColumn),
	).
		PlaceholderFormat(sq.Dollar).
		From(rep.EventTableName).
		LeftJoin(fmt.Sprintf("%s ON %s = %s",
			rep.EventDataTableName,
			rep.Col(rep.EventTableName, rep.EventIdColumn),
			rep.Col(rep.EventDataTableName, rep.EventDataEventIdColumn),
		)).
		LeftJoin(fmt.Sprintf("%s ON %s = %s",
			rep.EventStatusTableName,
			rep.Col(rep.EventTableName, rep.EventIdColumn),
			rep.Col(rep.EventStatusTableName, rep.EventStatusIdColumn),
		)).
		Where(sq.Eq{rep.Col(rep.EventTableName, rep.EventIdColumn): eventId}).
		ToSql()
	if err != nil {
		return nil, rep.NewInternalError(err)
	}

	row := r.db.DB().QueryRowContext(ctx, query, args...)

	//TODO: check how event status will be parsed
	event := &rmodel.Event{
		Data: &rmodel.EventData{},
	}
	err = row.Scan(
		&event.Id,

		&event.Data.CallerId,
		&event.Data.UtilizatorId,
		&event.Data.Latitude,
		&event.Data.Longtitude,
		&event.Data.Title,
		&event.Data.Description,
		&event.Data.TimeCalled,
		&event.Data.TimeUtilized,

		&event.Status,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, rep.NewNotFoundError(err)
		}
		return nil, rep.NewInternalError(err)
	}

	return event, nil
}

func (r *repo) GetCallerIdById(ctx context.Context, eventId uint64) (uint64, error) {
	query, args, err := sq.Select(rep.Col(rep.EventDataTableName, rep.EventDataCallerIdColumn)).
		PlaceholderFormat(sq.Dollar).
		From(rep.EventDataTableName).
		Where(sq.Eq{rep.Col(rep.EventDataTableName, rep.EventDataEventIdColumn): eventId}).
		ToSql()
	if err != nil {
		return 0, rep.NewInternalError(err)
	}

	var callerId uint64
	err = r.db.DB().QueryRowContext(ctx, query, args...).Scan(&callerId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, rep.NewNotFoundError(err)
		}
		return 0, rep.NewInternalError(err)
	}

	return callerId, nil
}
