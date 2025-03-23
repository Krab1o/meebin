package user

import (
	"context"
	"errors"
	"fmt"

	rmodel "github.com/Krab1o/meebin/internal/model/r_model"
	rep "github.com/Krab1o/meebin/internal/repository"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

// TODO: move function to rep layer

func (r *repo) GetById(ctx context.Context, userId uint64) (*rmodel.User, error) {
	query, args, err := sq.Select(
		rep.Col(rep.UserTableName, rep.UserEmailColumn),
		rep.Col(rep.UserTableName, rep.UserUsernameColumn),

		rep.Col(rep.DataTableName, rep.DataGivenNameColumn),
		rep.Col(rep.DataTableName, rep.DataSurnameColumn),
		rep.Col(rep.DataTableName, rep.DataPatronymicColumn),
		rep.Col(rep.DataTableName, rep.DataBirthDateColumn),
		rep.Col(rep.DataTableName, rep.DataCityColumn),

		rep.Col(rep.StatsTableName, rep.StatsReportCounterColumn),
		rep.Col(rep.StatsTableName, rep.StatsUtilizeCounterColumn),
		rep.Col(rep.StatsTableName, rep.StatsRatingColumn),
	).
		PlaceholderFormat(sq.Dollar).
		From(rep.UserTableName).
		LeftJoin(fmt.Sprintf("%s ON %s = %s",
			rep.DataTableName,
			rep.Col(rep.UserTableName, rep.UserIdColumn),
			rep.Col(rep.DataTableName, rep.DataIdUserColumn),
		)).
		LeftJoin(fmt.Sprintf("%s ON %s = %s",
			rep.StatsTableName,
			rep.Col(rep.UserTableName, rep.UserIdColumn),
			rep.Col(rep.StatsTableName, rep.StatsIdUserColumn),
		)).
		Where(sq.Eq{rep.Col(rep.UserTableName, rep.UserIdColumn): userId}).
		ToSql()
	if err != nil {
		return nil, rep.NewInternalError(err)
	}

	row := r.db.DB().QueryRowContext(ctx, query, args...)

	user := &rmodel.User{
		Creds: &rmodel.Creds{},
		Data:  &rmodel.PersonalData{},
		Stats: &rmodel.Stats{},
	}
	err = row.Scan(
		&user.Creds.Email,
		&user.Creds.Username,

		&user.Data.GivenName,
		&user.Data.Surname,
		&user.Data.Patronymic,
		&user.Data.Birthdate,
		&user.Data.City,

		&user.Stats.ReportCount,
		&user.Stats.UtilizeCount,
		&user.Stats.Rating,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, rep.NewNotFoundError(err)
		}
		return nil, rep.NewInternalError(err)
	}

	return user, nil
}
