package user

import (
	"context"

	rmodel "github.com/Krab1o/meebin/internal/model/r_model"
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) Add(ctx context.Context, tx pgx.Tx, user *rmodel.User) (uint64, error) {
	query, args, err := squirrel.Insert(repository.UserTableName).
		Columns(
			repository.UserUsernameColumn,
			repository.UserEmailColumn,
			repository.UserPasswordColumn,
		).
		Values(
			user.Creds.Username,
			user.Creds.Email,
			user.Creds.Password,
		).
		Suffix("RETURNING %s", repository.UserIdColumn).
		ToSql()
	if err != nil {
		return 0, err
	}

	var row pgx.Row
	if tx != nil {
		row = tx.QueryRow(ctx, query, args...)
	} else {
		row = r.db.QueryRow(ctx, query, args...)
	}

	var id uint64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
