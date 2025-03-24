package user

import (
	"context"

	rmodel "github.com/Krab1o/meebin/internal/model/user/r_model"
	rep "github.com/Krab1o/meebin/internal/repository"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) UpdateCreds(
	ctx context.Context,
	userId uint64,
	creds *rmodel.Creds,
) error {
	builder := sq.Update(rep.UserTableName).
		PlaceholderFormat(sq.Dollar)
	if creds.Email != "" {
		builder = builder.Set(rep.UserEmailColumn, creds.Email)
	}
	if creds.HashedPassword != "" {
		builder = builder.Set(
			rep.UserPasswordColumn,
			creds.HashedPassword,
		)
	}
	if creds.Username != "" {
		builder = builder.Set(rep.UserUsernameColumn, creds.Username)
	}
	builder = builder.Where(sq.Eq{rep.UserIdColumn: userId})
	query, args, err := builder.ToSql()
	if err != nil {
		return rep.NewInternalError(err)
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)

	if err != nil {
		return rep.NewInternalError(err)
	}

	return nil
}

func (r *repo) UpdateStats(
	ctx context.Context,
	userId uint64,
	stats *rmodel.Stats,
) error {
	builder := sq.Update(rep.StatsTableName).
		PlaceholderFormat(sq.Dollar)
	if stats.UtilizeCount != 0 {
		builder = builder.Set(
			rep.StatsUtilizeCounterColumn,
			stats.UtilizeCount,
		)
	}
	if stats.ReportCount != 0 {
		builder = builder.Set(
			rep.StatsReportCounterColumn,
			stats.ReportCount,
		)
	}
	if stats.Rating != 0.0 {
		builder = builder.Set(rep.StatsRatingColumn, stats.Rating)
	}
	builder = builder.Where(sq.Eq{rep.UserIdColumn: userId})
	query, args, err := builder.ToSql()
	if err != nil {
		return rep.NewInternalError(err)
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)

	if err != nil {
		return rep.NewInternalError(err)
	}
	return nil
}

// TODO: refactor to passing map[string]string of values
func (r *repo) UpdatePersonalData(
	ctx context.Context,
	userId uint64,
	data *rmodel.PersonalData,
) error {
	builder := sq.Update(rep.UserDataTableName).
		PlaceholderFormat(sq.Dollar)
	if data.GivenName != "" {
		builder = builder.Set(rep.UserDataGivenNameColumn, data.GivenName)
	}
	if data.Surname != "" {
		builder = builder.Set(rep.UserDataSurnameColumn, data.Surname)
	}
	if data.Patronymic != "" {
		builder = builder.Set(rep.UserDataPatronymicColumn, data.Patronymic)
	}
	if data.City != "" {
		builder = builder.Set(rep.UserDataCityColumn, data.City)
	}
	if data.Birthdate.IsZero() {
		builder = builder.Set(rep.UserDataBirthDateColumn, data.Birthdate)
	}

	builder = builder.Where(sq.Eq{rep.UserIdColumn: userId})
	query, args, err := builder.ToSql()
	if err != nil {
		return rep.NewInternalError(err)
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)

	if err != nil {
		return rep.NewInternalError(err)
	}
	return nil
}
