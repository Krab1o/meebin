package session

import (
	"context"

	"github.com/Krab1o/meebin/internal/repository"
	rmodel "github.com/Krab1o/meebin/internal/struct/r_model"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) AddSession(ctx context.Context, tx pgx.Tx, session *rmodel.Session) error {
	sessionTableQuery, sessionTableArgs, err := squirrel.
		Insert(repository.SessionTableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			repository.SessionIdUserColumn,
			repository.SessionRefreshTokenColumn,
			repository.SessionExpirationTimeColumn,
		).
		Values(
			session.UserId,
			session.RefreshToken,
			session.ExpirationTime,
		).
		ToSql()
	if err != nil {
		return err
	}
	if tx != nil {
		_, err = tx.Exec(ctx, sessionTableQuery, sessionTableArgs...)
	} else {
		_, err = r.db.Exec(ctx, sessionTableQuery, sessionTableArgs...)
	}
	if err != nil {
		return err
	}
	return nil
}
