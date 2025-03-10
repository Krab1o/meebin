package session

import (
	"context"
	"fmt"

	rmodel "github.com/Krab1o/meebin/internal/model/r_model"
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) AddSession(ctx context.Context, tx pgx.Tx, session *rmodel.Session) (uint64, error) {
	sessionTableQuery, sessionTableArgs, err := squirrel.
		Insert(repository.SessionTableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			repository.SessionIdUserColumn,
			repository.SessionExpirationTimeColumn,
		).
		Values(
			session.UserId,
			session.ExpirationTime,
		).
		Suffix(
			fmt.Sprintf("RETURNING %s", repository.SessionIdColumn),
		).
		ToSql()
	if err != nil {
		return 0, repository.NewInternalError(err)
	}
	var row pgx.Row
	if tx != nil {
		row = tx.QueryRow(ctx, sessionTableQuery, sessionTableArgs...)
	} else {
		row = r.db.QueryRow(ctx, sessionTableQuery, sessionTableArgs...)
	}
	var sessionId uint64
	err = row.Scan(&sessionId)
	if err != nil {
		return 0, repository.NewInternalError(err)
	}
	return sessionId, nil
}
