package session

import (
	"context"
	"errors"
	"log"

	rmodel "github.com/Krab1o/meebin/internal/model/r_model"
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) FindSession(
	ctx context.Context,
	tx pgx.Tx,
	sessionID uint64,
) (*rmodel.Session, error) {
	query, args, err := squirrel.Select(
		repository.SessionIdColumn,
		repository.SessionIdUserColumn,
		repository.SessionExpirationTimeColumn,
	).
		PlaceholderFormat(squirrel.Dollar).
		From(repository.SessionTableName).
		Where(squirrel.Eq{repository.SessionIdColumn: sessionID}).
		ToSql()
	log.Println(query)
	if err != nil {
		return nil, repository.NewInternalError(err)
	}

	var row pgx.Row
	if tx != nil {
		row = tx.QueryRow(ctx, query, args...)
	} else {
		row = r.db.QueryRow(ctx, query, args...)
	}

	repoSession := &rmodel.Session{}
	err = row.Scan(
		&repoSession.SessionId,
		&repoSession.UserId,
		&repoSession.ExpirationTime,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.NewNotFoundError(err)
		}
		return nil, repository.NewInternalError(err)
	}
	return repoSession, nil
}
