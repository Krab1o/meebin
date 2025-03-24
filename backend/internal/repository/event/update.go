package event

import (
	"context"

	"github.com/Krab1o/meebin/internal/model"
	rmodel "github.com/Krab1o/meebin/internal/model/event/r_model"
	rep "github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
)

func (r *repo) UpdateStatus(ctx context.Context, eventId uint64, status model.EventStatus) (*rmodel.Event, error) {
	builder := sq.Update(rep.EventStatusColumn).PlaceholderFormat(squirrel.Dollar)
}
