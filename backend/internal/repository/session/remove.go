package session

import (
	"context"

	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) DeleteSession(ctx context.Context, tx pgx.Tx, sessionId uint64) error {
	query, args, err := squirrel.Delete(repository.SessionTableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{repository.SessionIdColumn: sessionId}).
		ToSql()
	if err != nil {
		return repository.NewInternalError(err)
	}

	if tx != nil {
		_, err = tx.Exec(ctx, query, args...)
	} else {
		_, err = r.db.Exec(ctx, query, args...)
	}
	if err != nil {
		return repository.NewInternalError(err)
	}
	return nil
}
