package user

import (
	"github.com/Krab1o/meebin/internal/model"
	"github.com/Krab1o/meebin/internal/model/dto"
	smodel "github.com/Krab1o/meebin/internal/model/s_model"
)

// TODO: add role to user
func UserServiceToDTO(user *smodel.User) *dto.User {
	dtoCreds := CredsServiceToDTO(user.Creds)
	dtoData := DataServiceToDTO(user.Data)
	dtoStats := StatsServiceToDTO(user.Stats)
	dtoRoles := RolesServiceToDTO(user.Roles)
	return &dto.User{
		Id:    user.Id,
		Creds: dtoCreds,
		Data:  dtoData,
		Stats: dtoStats,
		Roles: dtoRoles,
	}
}

func RolesServiceToDTO(roles []model.Role) []model.Role {
	if roles == nil {
		return nil
	}
	copiedRoles := make([]model.Role, len(roles))
	copy(copiedRoles, roles)

	return copiedRoles
}

func CredsServiceToDTO(creds *smodel.Creds) *dto.Creds {
	if creds == nil {
		return nil
	}
	return &dto.Creds{
		Username: creds.Username,
		Email:    creds.Email,
	}
}

func DataServiceToDTO(data *smodel.PersonalData) *dto.PersonalData {
	if data == nil {
		return nil
	}
	return &dto.PersonalData{
		GivenName:  data.GivenName,
		Surname:    data.Surname,
		Patronymic: data.Patronymic,
		Birthdate:  data.Birthdate,
		City:       data.City,
	}
}

func StatsServiceToDTO(stats *smodel.Stats) *dto.Stats {
	if stats == nil {
		return nil
	}
	return &dto.Stats{
		UtilizeCount: stats.UtilizeCount,
		ReportCount:  stats.ReportCount,
		Rating:       stats.Rating,
	}
}
