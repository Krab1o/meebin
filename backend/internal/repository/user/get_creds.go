package user

import (
	"context"
	"errors"

	rmodel "github.com/Krab1o/meebin/internal/model/user/r_model"
	rep "github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) GetCredsByEmail(
	ctx context.Context,
	email string,
) (*rmodel.User, error) {
	query, args, err := squirrel.
		Select(
			rep.UserIdColumn,
			rep.UserUsernameColumn,
			rep.UserEmailColumn,
			rep.UserPasswordColumn,
		).
		PlaceholderFormat(squirrel.Dollar).
		From(rep.UserTableName).
		Where(squirrel.Eq{rep.UserEmailColumn: email}).
		ToSql()
	if err != nil {
		return nil, rep.NewInternalError(err)
	}

	row := r.db.DB().QueryRowContext(ctx, query, args...)

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
			return nil, rep.NewNotFoundError(err)
		}
		return nil, rep.NewInternalError(err)
	}
	return user, nil
}
