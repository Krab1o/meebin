package session

import (
	"context"
	"errors"
	"log"

	rmodel "github.com/Krab1o/meebin/internal/model/user/r_model"
	rep "github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) FindSessionById(
	ctx context.Context,
	sessionID uint64,
) (*rmodel.Session, error) {
	query, args, err := squirrel.Select(
		rep.SessionColumnId,
		rep.SessionColumnIdUser,
		rep.SessionColumnExpirationTime,
	).
		PlaceholderFormat(squirrel.Dollar).
		From(rep.SessionTableName).
		Where(squirrel.Eq{rep.SessionColumnId: sessionID}).
		ToSql()
	log.Println(query)
	if err != nil {
		return nil, rep.NewInternalError(err)
	}

	row := r.db.DB().QueryRowContext(ctx, query, args...)

	repoSession := &rmodel.Session{}
	err = row.Scan(
		&repoSession.SessionId,
		&repoSession.UserId,
		&repoSession.ExpirationTime,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, rep.NewNotFoundError(err)
		}
		return nil, rep.NewInternalError(err)
	}
	return repoSession, nil
}
