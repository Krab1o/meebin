package user

import (
	"github.com/Krab1o/meebin/internal/converter"
	rmodel "github.com/Krab1o/meebin/internal/model/user/r_model"
	smodel "github.com/Krab1o/meebin/internal/model/user/s_model"
)

func UserRepoToService(user *rmodel.User) *smodel.User {
	serviceCreds := CredsRepoToService(user.Creds)
	serviceData := DataRepoToService(user.Data)
	serviceStats := StatsRepoToService(user.Stats)
	serviceRoles := converter.ConvertRoles(user.Roles)
	return &smodel.User{
		Id:       user.Id,
		Roles:    serviceRoles,
		Creds:    serviceCreds,
		Data:     serviceData,
		Stats:    serviceStats,
		Sessions: nil,
	}
}

func CredsRepoToService(creds *rmodel.Creds) *smodel.Creds {
	if creds == nil {
		return nil
	}
	return &smodel.Creds{
		Username: creds.Username,
		Email:    creds.Email,
	}
}

func DataRepoToService(data *rmodel.PersonalData) *smodel.PersonalData {
	if data == nil {
		return nil
	}
	return &smodel.PersonalData{
		GivenName:  data.GivenName,
		Surname:    data.Surname,
		Patronymic: data.Patronymic,
		Birthdate:  data.Birthdate,
		City:       data.City,
	}
}

func StatsRepoToService(stats *rmodel.Stats) *smodel.Stats {
	if stats == nil {
		return nil
	}
	return &smodel.Stats{
		UtilizeCount: stats.UtilizeCount,
		ReportCount:  stats.ReportCount,
		Rating:       stats.Rating,
	}
}
