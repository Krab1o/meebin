package converter

import (
	"github.com/Krab1o/meebin/internal/model"
	rmodel "github.com/Krab1o/meebin/internal/model/r_model"
	smodel "github.com/Krab1o/meebin/internal/model/s_model"
)

func UserRepoToService(user *rmodel.User) *smodel.User {

	return &smodel.User{
		Id: user.Id,
	}
}

func RolesRepoToService(roles []model.Role)

func CredsRepoToService(creds *rmodel.Creds) *smodel.Creds {
	return &smodel.Creds{}
}

func DataRepoToService(data *rmodel.PersonalData) *smodel.PersonalData {
	return &smodel.PersonalData{}
}

func StatsRepoToService(stats *rmodel.Stats) *smodel.Stats {
	return &smodel.Stats{}
}
