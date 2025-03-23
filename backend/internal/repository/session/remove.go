package session

import (
	"context"

	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
)

func (r *repo) DeleteSession(ctx context.Context, sessionId uint64) error {
	query, args, err := squirrel.Delete(repository.SessionTableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{repository.SessionIdColumn: sessionId}).
		ToSql()
	if err != nil {
		return repository.NewInternalError(err)
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return repository.NewInternalError(err)
	}
	return nil
}
