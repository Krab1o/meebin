package user

import (
	"context"
	"errors"

	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) DeleteById(ctx context.Context, tx pgx.Tx, userId uint64) error {
	query, args, err := squirrel.Delete(repository.UserTableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{repository.UserIdColumn: userId}).
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
		if errors.Is(err, pgx.ErrNoRows) {
			return repository.NewNotFoundError(err)
		}
		return repository.NewInternalError(err)
	}
	return nil
}
