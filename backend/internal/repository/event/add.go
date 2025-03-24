package event

import (
	"context"
	"errors"
	"fmt"

	rmodel "github.com/Krab1o/meebin/internal/model/event/r_model"
	rep "github.com/Krab1o/meebin/internal/repository"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *repo) addEvent(ctx context.Context, newEvent *rmodel.Event) (uint64, error) {
	query, args, err := sq.Insert(rep.EventDataTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			rep.EventStatusColumn,
		).
		Values(
			newEvent.Status,
		).
		Suffix(
			fmt.Sprintf("RETURNING %s", rep.SessionIdColumn),
		).
		ToSql()
	if err != nil {
		return 0, rep.NewInternalError(err)
	}

	var eventId uint64
	err = r.db.DB().QueryRowContext(ctx, query, args...).Scan(&eventId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case rep.SQLCodeDuplicate:
				return 0, rep.NewDuplicateError(err)
			}
		}
		return 0, rep.NewInternalError(err)
	}

	return eventId, nil
}

func (r *repo) addEventData(ctx context.Context, newEventData *rmodel.EventData) error {
	query, args, err := sq.Insert(rep.EventDataTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			rep.EventDataLatitudeColumn,
			rep.EventDataLongtitudeColumn,
			rep.EventDataTitleColumn,
			rep.EventDataDescriptionColumn,
			rep.EventDataTimeCalledColumn,
			rep.EventDataTimeCleanedColumn,
		).Values(
		newEventData.Latitude,
		newEventData.Longtitude,
		newEventData.Title,
		newEventData.Description,
		newEventData.TimeCalled,
		newEventData.TimeCleaned,
	).ToSql()
	if err != nil {
		return rep.NewInternalError(err)
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return rep.NewInternalError(err)
	}
	return nil
}

// TODO: decide how to check for nil (and if I need to check at all)
func (r *repo) Add(ctx context.Context, newEvent *rmodel.Event) (uint64, error) {
	eventId, err := r.addEvent(ctx, newEvent)
	if err != nil {
		return 0, err
	}
	err = r.addEventData(ctx, newEvent.Data)
	if err != nil {
		return 0, err
	}
	return eventId, nil
}
