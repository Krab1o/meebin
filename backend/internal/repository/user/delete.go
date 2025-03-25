package user

import (
	"context"
	"errors"

	rep "github.com/Krab1o/meebin/internal/repository"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) DeleteUserById(ctx context.Context, userId uint64) error {
	query, args, err := sq.Delete(rep.UserTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{rep.UserIdColumn: userId}).
		ToSql()
	if err != nil {
		return rep.NewInternalError(err)
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return rep.NewNotFoundError(err)
		}
		return rep.NewInternalError(err)
	}
	return nil
}
