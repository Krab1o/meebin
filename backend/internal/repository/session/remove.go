package session

import (
	"context"

	rep "github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
)

func (r *repo) DeleteSessionById(ctx context.Context, sessionId uint64) error {
	query, args, err := squirrel.Delete(rep.SessionTableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{rep.SessionColumnId: sessionId}).
		ToSql()
	if err != nil {
		return rep.NewInternalError(err)
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return rep.NewInternalError(err)
	}
	return nil
}
