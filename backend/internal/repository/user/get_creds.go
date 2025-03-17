package user

import (
	"context"
	"errors"

	rmodel "github.com/Krab1o/meebin/internal/model/r_model"
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) GetCredsByEmail(
	ctx context.Context,
	tx pgx.Tx,
	email string,
) (*rmodel.User, error) {
	query, args, err := squirrel.
		Select(
			repository.UserIdColumn,
			repository.UserUsernameColumn,
			repository.UserEmailColumn,
			repository.UserPasswordColumn,
		).
		PlaceholderFormat(squirrel.Dollar).
		From(repository.UserTableName).
		Where(squirrel.Eq{repository.UserEmailColumn: email}).
		ToSql()
	if err != nil {
		return nil, repository.NewInternalError(err)
	}
	var row pgx.Row
	if tx != nil {
		row = tx.QueryRow(ctx, query, args...)
	} else {
		row = r.db.QueryRow(ctx, query, args...)
	}
	user := &rmodel.User{
		Creds: &rmodel.Creds{},
	}
	err = row.Scan(
		&user.Id,
		&user.Creds.Username,
		&user.Creds.Email,
		&user.Creds.HashedPassword,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.NewNotFoundError(err)
		}
		return nil, repository.NewInternalError(err)
	}
	return user, nil
}
