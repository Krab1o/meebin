package event

import (
	"context"
	"fmt"

	rmodel "github.com/Krab1o/meebin/internal/model/event/r_model"
	rep "github.com/Krab1o/meebin/internal/repository"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) List(ctx context.Context) ([]rmodel.Event, error) {
	query, args, err := sq.Select(
		rep.Col(rep.EventTableName, rep.EventIdColumn),

		rep.Col(rep.EventTableName, rep.EventDataCallerIdColumn),
		rep.Col(rep.EventTableName, rep.EventDataUtilizatorIdColumn),
		rep.Col(rep.EventDataTableName, rep.EventDataLatitudeColumn),
		rep.Col(rep.EventDataTableName, rep.EventDataLongtitudeColumn),
		rep.Col(rep.EventDataTableName, rep.EventDataTitleColumn),
		rep.Col(rep.EventDataTableName, rep.EventDataDescriptionColumn),
		rep.Col(rep.EventDataTableName, rep.EventDataTimeCalledColumn),
		rep.Col(rep.EventDataTableName, rep.EventDataTimeCleanedColumn),

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
		ToSql()
	if err != nil {
		return nil, rep.NewInternalError(err)
	}

	var rows pgx.Rows
	rows, err = r.db.DB().QueryContext(ctx, query, args...)

	if err != nil {
		return nil, rep.NewInternalError(err)
	}

	list := []rmodel.Event{}
	var event *rmodel.Event

	//TODO: check how event status will be parsed
	for rows.Next() {
		event = &rmodel.Event{
			Data: &rmodel.EventData{},
		}
		err = rows.Scan(
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
			return nil, rep.NewInternalError(err)
		}
		list = append(list, *event)
	}
	return list, nil
}
