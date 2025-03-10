package user

import (
	"context"
	"errors"
	"fmt"

	rmodel "github.com/Krab1o/meebin/internal/model/r_model"
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// TODO: cache roles to avoid queries
// TODO: add multiple role support
func (r *repo) AddUser(
	ctx context.Context,
	tx pgx.Tx,
	user *rmodel.User,
	roleId uint64,
) (uint64, error) {
	userTableQuery, userTableArgs, err := squirrel.Insert(repository.UserTableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			repository.UserUsernameColumn,
			repository.UserEmailColumn,
			repository.UserPasswordColumn,
		).
		Values(
			user.Creds.Username,
			user.Creds.Email,
			user.Creds.HashedPassword,
		).
		Suffix(fmt.Sprintf("RETURNING %s", repository.UserIdColumn)).
		ToSql()
	if err != nil {
		return 0, repository.NewInternalError(err)
	}

	// TODO: Refactor with transaction manager or add query interface
	// Check how should transaction manager be used and how it works
	var row pgx.Row
	if tx != nil {
		row = tx.QueryRow(ctx, userTableQuery, userTableArgs...)
	} else {
		row = r.db.QueryRow(ctx, userTableQuery, userTableArgs...)
	}

	var userId uint64
	err = row.Scan(&userId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case repository.SQLCodeDuplicate:
				return 0, repository.NewDuplicateError(err)
			}
		}
		return 0, repository.NewInternalError(err)
	}

	dataTableQuery, dataTableArgs, err := squirrel.Insert(repository.DataTableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			repository.DataIdUserColumn,
			repository.DataGivenNameColumn,
			repository.DataSurnameColumn,
			repository.DataPatronymicColumn,
			repository.DataCityColumn,
			repository.DataBirthDateColumn,
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
		return 0, repository.NewInternalError(err)
	}

	if tx != nil {
		_, err = tx.Exec(ctx, dataTableQuery, dataTableArgs...)
	} else {
		_, err = r.db.Exec(ctx, dataTableQuery, dataTableArgs...)
	}
	if err != nil {
		return 0, repository.NewInternalError(err)
	}

	statsTableQuery, statsTableArgs, err := squirrel.Insert(repository.StatsTableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			repository.StatsIdUserColumn,
			repository.StatsUtilizeCounterColumn,
			repository.StatsReportCounterColumn,
			repository.StatsRatingColumn,
		).
		Values(
			userId,
			user.Stats.UtilizeCount,
			user.Stats.ReportCount,
			user.Stats.Rating,
		).ToSql()
	if err != nil {
		return 0, repository.NewInternalError(err)
	}

	if tx != nil {
		_, err = tx.Exec(ctx, statsTableQuery, statsTableArgs...)
	} else {
		_, err = r.db.Exec(ctx, statsTableQuery, statsTableArgs...)
	}
	if err != nil {
		return 0, repository.NewInternalError(err)
	}

	userRoleQuery, userRoleArgs, err := squirrel.Insert(repository.UserRoleTableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			repository.UserRoleIdRoleColumn,
			repository.UserRoleIdUserColumn,
		).Values(
		roleId,
		userId,
	).ToSql()
	if err != nil {
		return 0, repository.NewInternalError(err)
	}
	if tx != nil {
		_, err = tx.Exec(ctx, userRoleQuery, userRoleArgs...)
	} else {
		_, err = r.db.Exec(ctx, userRoleQuery, userRoleArgs...)
	}
	if err != nil {
		return 0, repository.NewInternalError(err)
	}
	return userId, nil
}
