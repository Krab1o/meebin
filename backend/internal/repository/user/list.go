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
		rep.Col(rep.UserTableName, rep.UserColumnEmail),
		rep.Col(rep.UserTableName, rep.UserColumnUsername),

		rep.Col(rep.UserDataTableName, rep.UserDataColumnGivenName),
		rep.Col(rep.UserDataTableName, rep.UserDataColumnSurname),
		rep.Col(rep.UserDataTableName, rep.UserDataColumnPatronymic),
		rep.Col(rep.UserDataTableName, rep.UserDataColumnBirthDate),
		rep.Col(rep.UserDataTableName, rep.UserDataColumnCity),

		rep.Col(rep.StatsTableName, rep.StatsColumnReportCounter),
		rep.Col(rep.StatsTableName, rep.StatsColumnUtilizeCounter),
		rep.Col(rep.StatsTableName, rep.StatsColumnRating),
	).
		PlaceholderFormat(sq.Dollar).
		From(rep.UserTableName).
		LeftJoin(fmt.Sprintf("%s ON %s = %s",
			rep.UserDataTableName,
			rep.Col(rep.UserTableName, rep.UserColumnId),
			rep.Col(rep.UserDataTableName, rep.UserDataColumnIdUser),
		)).
		LeftJoin(fmt.Sprintf("%s ON %s = %s",
			rep.StatsTableName,
			rep.Col(rep.UserTableName, rep.UserColumnId),
			rep.Col(rep.StatsTableName, rep.StatsColumnIdUser),
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
