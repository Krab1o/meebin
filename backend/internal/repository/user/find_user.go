package user

import (
	"context"

	"github.com/Krab1o/meebin/internal/repository"
	rmodel "github.com/Krab1o/meebin/internal/struct/r_model"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

// TODO: add returning values for user
func (r *repo) FindUser(ctx context.Context, tx pgx.Tx, creds *rmodel.Creds) (uint64, error) {
	query, args, err := squirrel.
		Select(repository.UserIdColumn).
		PlaceholderFormat(squirrel.Dollar).
		From(repository.UserTableName).
		Where(squirrel.Eq{repository.UserEmailColumn: creds.Email}).
		ToSql()
	if err != nil {
		return 0, repository.NewInternalError(err)
	}
	var row pgx.Row
	if tx != nil {
		row = tx.QueryRow(ctx, query, args...)
	} else {
		row = r.db.QueryRow(ctx, query, args...)
	}
	var userId uint64
	err = row.Scan(&userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, repository.NewNotFoundError(err)
		}
		return 0, repository.NewInternalError(err)
	}
	return userId, nil
}
