package converter

import (
	"github.com/Krab1o/meebin/internal/model"
	rmodel "github.com/Krab1o/meebin/internal/model/r_model"
	smodel "github.com/Krab1o/meebin/internal/model/s_model"
)

func UserServiceToRepo(user *smodel.User) *rmodel.User {
	repoCreds := CredsServiceToRepo(user.Creds)
	repoStats := StatsServiceToRepo(user.Stats)
	repoData := PersonalDataServiceToRepo(user.Data)
	repoRoles := RolesServiceToRepo(user.Roles)
	return &rmodel.User{
		Id:    user.Id,
		Roles: repoRoles,
		Creds: repoCreds,
		Stats: repoStats,
		Data:  repoData,
	}
}

func RolesServiceToRepo(roles []model.Role) []model.Role {
	if roles == nil {
		return nil
	}

	copiedRoles := make([]model.Role, len(roles))
	copy(copiedRoles, roles)

	return copiedRoles
}

func StatsServiceToRepo(stats *smodel.Stats) *rmodel.Stats {
	return &rmodel.Stats{
		UtilizeCount: stats.UtilizeCount,
		ReportCount:  stats.ReportCount,
		Rating:       stats.Rating,
	}
}

func PersonalDataServiceToRepo(data *smodel.PersonalData) *rmodel.PersonalData {
	return &rmodel.PersonalData{
		GivenName:  data.GivenName,
		Surname:    data.Surname,
		Patronymic: data.Patronymic,
		City:       data.City,
		Birthdate:  data.Birthdate,
	}
}

func CredsServiceToRepo(creds *smodel.Creds) *rmodel.Creds {
	return &rmodel.Creds{
		Username:       creds.Username,
		Email:          creds.Email,
		HashedPassword: creds.Password,
	}
}
