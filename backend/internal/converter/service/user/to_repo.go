package user

import (
	"github.com/Krab1o/meebin/internal/converter"
	rmodel "github.com/Krab1o/meebin/internal/model/user/r_model"
	smodel "github.com/Krab1o/meebin/internal/model/user/s_model"
)

func UserServiceToRepo(user *smodel.User) *rmodel.User {
	repoCreds := CredsServiceToRepo(user.Creds)
	repoStats := StatsServiceToRepo(user.Stats)
	repoData := PersonalDataServiceToRepo(user.Data)
	repoRoles := converter.ConvertRoles(user.Roles)
	return &rmodel.User{
		Id:    user.Id,
		Roles: repoRoles,
		Creds: repoCreds,
		Stats: repoStats,
		Data:  repoData,
	}
}

func StatsServiceToRepo(stats *smodel.Stats) *rmodel.Stats {
	if stats == nil {
		return nil
	}
	return &rmodel.Stats{
		UtilizeCount: stats.UtilizeCount,
		ReportCount:  stats.ReportCount,
		Rating:       stats.Rating,
	}
}

func PersonalDataServiceToRepo(data *smodel.PersonalData) *rmodel.PersonalData {
	if data == nil {
		return nil
	}
	return &rmodel.PersonalData{
		GivenName:  data.GivenName,
		Surname:    data.Surname,
		Patronymic: data.Patronymic,
		City:       data.City,
		Birthdate:  data.Birthdate,
	}
}

func CredsServiceToRepo(creds *smodel.Creds) *rmodel.Creds {
	if creds == nil {
		return nil
	}
	return &rmodel.Creds{
		Username:       creds.Username,
		Email:          creds.Email,
		HashedPassword: creds.Password,
	}
}
