package user

import (
	"context"
	"fmt"

	rmodel "github.com/Krab1o/meebin/internal/model/user/r_model"
	rep "github.com/Krab1o/meebin/internal/repository"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) List(ctx context.Context) ([]rmodel.User, error) {
	query, args, err := sq.Select(
		rep.Col(rep.UserTableName, rep.UserEmailColumn),
		rep.Col(rep.UserTableName, rep.UserUsernameColumn),

		rep.Col(rep.UserDataTableName, rep.UserDataGivenNameColumn),
		rep.Col(rep.UserDataTableName, rep.UserDataSurnameColumn),
		rep.Col(rep.UserDataTableName, rep.UserDataPatronymicColumn),
		rep.Col(rep.UserDataTableName, rep.UserDataBirthDateColumn),
		rep.Col(rep.UserDataTableName, rep.UserDataCityColumn),

		rep.Col(rep.StatsTableName, rep.StatsReportCounterColumn),
		rep.Col(rep.StatsTableName, rep.StatsUtilizeCounterColumn),
		rep.Col(rep.StatsTableName, rep.StatsRatingColumn),
	).
		PlaceholderFormat(sq.Dollar).
		From(rep.UserTableName).
		LeftJoin(fmt.Sprintf("%s ON %s = %s",
			rep.UserDataTableName,
			rep.Col(rep.UserTableName, rep.UserIdColumn),
			rep.Col(rep.UserDataTableName, rep.UserDataIdUserColumn),
		)).
		LeftJoin(fmt.Sprintf("%s ON %s = %s",
			rep.StatsTableName,
			rep.Col(rep.UserTableName, rep.UserIdColumn),
			rep.Col(rep.StatsTableName, rep.StatsIdUserColumn),
		)).
		ToSql()
	if err != nil {
		return nil, rep.NewInternalError(err)
	}

	var rows pgx.Rows
	rows, err = r.db.DB().QueryContext(ctx, query, args...)

	if err != nil {
		return nil, rep.NewInternalError(err)
	}

	list := []rmodel.User{}
	var user *rmodel.User

	for rows.Next() {
		user = &rmodel.User{
			Creds: &rmodel.Creds{},
			Data:  &rmodel.PersonalData{},
			Stats: &rmodel.Stats{},
		}
		err = rows.Scan(
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
			return nil, rep.NewInternalError(err)
		}
		list = append(list, *user)
	}
	return list, nil
}
