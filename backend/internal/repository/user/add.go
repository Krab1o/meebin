package user

import (
	"context"
	"errors"
	"fmt"

	rmodel "github.com/Krab1o/meebin/internal/model/user/r_model"
	rep "github.com/Krab1o/meebin/internal/repository"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
)

// TODO: cache roles to avoid queries
// TODO: add multiple role support
// TODO: refactor splitting to function
func (r *repo) AddUser(
	ctx context.Context,
	user *rmodel.User,
	roleId uint64,
) (uint64, error) {
	userTableQuery, userTableArgs, err := sq.Insert(rep.UserTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			rep.UserUsernameColumn,
			rep.UserEmailColumn,
			rep.UserPasswordColumn,
		).
		Values(
			user.Creds.Username,
			user.Creds.Email,
			user.Creds.HashedPassword,
		).
		Suffix(fmt.Sprintf("RETURNING %s", rep.UserIdColumn)).
		ToSql()
	if err != nil {
		return 0, rep.NewInternalError(err)
	}

	row := r.db.DB().QueryRowContext(ctx, userTableQuery, userTableArgs...)

	var userId uint64
	err = row.Scan(&userId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case rep.SQLCodeDuplicate:
				return 0, rep.NewDuplicateError(err)
			}
		}
		return 0, rep.NewInternalError(err)
	}

	dataTableQuery, dataTableArgs, err := sq.Insert(rep.UserDataTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			rep.UserDataIdUserColumn,
			rep.UserDataGivenNameColumn,
			rep.UserDataSurnameColumn,
			rep.UserDataPatronymicColumn,
			rep.UserDataCityColumn,
			rep.UserDataBirthDateColumn,
		).
		Values(
			userId,
			user.Data.GivenName,
			user.Data.Surname,
			user.Data.Patronymic,
			user.Data.City,
			user.Data.Birthdate,
		).ToSql()
	if err != nil {
		return 0, rep.NewInternalError(err)
	}

	_, err = r.db.DB().ExecContext(ctx, dataTableQuery, dataTableArgs...)

	if err != nil {
		return 0, rep.NewInternalError(err)
	}

	statsTableQuery, statsTableArgs, err := sq.Insert(rep.StatsTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			rep.StatsIdUserColumn,
			rep.StatsUtilizeCounterColumn,
			rep.StatsReportCounterColumn,
			rep.StatsRatingColumn,
		).
		Values(
			userId,
			user.Stats.UtilizeCount,
			user.Stats.ReportCount,
			user.Stats.Rating,
		).ToSql()
	if err != nil {
		return 0, rep.NewInternalError(err)
	}

	_, err = r.db.DB().ExecContext(ctx, statsTableQuery, statsTableArgs...)

	if err != nil {
		return 0, rep.NewInternalError(err)
	}

	userRoleQuery, userRoleArgs, err := sq.Insert(rep.UserRoleTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			rep.UserRoleIdRoleColumn,
			rep.UserRoleIdUserColumn,
		).Values(
		roleId,
		userId,
	).ToSql()
	if err != nil {
		return 0, rep.NewInternalError(err)
	}

	_, err = r.db.DB().ExecContext(ctx, userRoleQuery, userRoleArgs...)

	if err != nil {
		return 0, rep.NewInternalError(err)
	}
	return userId, nil
}
