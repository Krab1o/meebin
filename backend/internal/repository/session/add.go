package session

import (
	"context"
	"fmt"

	rmodel "github.com/Krab1o/meebin/internal/model/user/r_model"
	rep "github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
)

func (r *repo) Add(ctx context.Context, session *rmodel.Session) (uint64, error) {
	sessionTableQuery, sessionTableArgs, err := squirrel.
		Insert(rep.SessionTableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			rep.SessionIdUserColumn,
			rep.SessionExpirationTimeColumn,
		).
		Values(
			session.UserId,
			session.ExpirationTime,
		).
		Suffix(
			fmt.Sprintf("RETURNING %s", rep.SessionIdColumn),
		).
		ToSql()
	if err != nil {
		return 0, rep.NewInternalError(err)
	}
	row := r.db.DB().QueryRowContext(ctx, sessionTableQuery, sessionTableArgs...)

	var sessionId uint64
	err = row.Scan(&sessionId)
	if err != nil {
		return 0, rep.NewInternalError(err)
	}
	return sessionId, nil
}
