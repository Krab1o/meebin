package user

import (
	"context"
	"errors"

	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) DeleteById(ctx context.Context, userId uint64) error {
	query, args, err := squirrel.Delete(repository.UserTableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{repository.UserIdColumn: userId}).
		ToSql()
	if err != nil {
		return repository.NewInternalError(err)
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return repository.NewNotFoundError(err)
		}
		return repository.NewInternalError(err)
	}
	return nil
}
