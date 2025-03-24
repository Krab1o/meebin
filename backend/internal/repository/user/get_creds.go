package user

import (
	"context"
	"errors"

	rmodel "github.com/Krab1o/meebin/internal/model/user/r_model"
	rep "github.com/Krab1o/meebin/internal/repository"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) GetCredsByEmail(
	ctx context.Context,
	email string,
) (*rmodel.User, error) {
	query, args, err := sq.
		Select(
			rep.UserIdColumn,
			rep.UserUsernameColumn,
			rep.UserEmailColumn,
			rep.UserPasswordColumn,
		).
		PlaceholderFormat(sq.Dollar).
		From(rep.UserTableName).
		Where(sq.Eq{rep.UserEmailColumn: email}).
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
