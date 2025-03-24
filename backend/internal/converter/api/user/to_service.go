package user

import (
	"github.com/Krab1o/meebin/internal/converter"
	"github.com/Krab1o/meebin/internal/model/user/dto"
	smodel "github.com/Krab1o/meebin/internal/model/user/s_model"
)

func UserDTOToService(user *dto.BaseUser) *smodel.User {
	dtoCreds := CredsDTOToService(user.Creds)
	dtoData := DataDTOToService(user.Data)
	dtoStats := StatsDTOToService(user.Stats)
	dtoRoles := converter.ConvertRoles(user.Roles)
	return &smodel.User{
		Id:    user.Id,
		Roles: dtoRoles,
		Creds: dtoCreds,
		Data:  dtoData,
		Stats: dtoStats,
	}
}

func CredsDTOToService(creds *dto.Creds) *smodel.Creds {
	if creds == nil {
		return nil
	}
	return &smodel.Creds{
		Username: creds.Username,
		Email:    creds.Email,
		Password: creds.Password,
	}
}

func DataDTOToService(data *dto.PersonalData) *smodel.PersonalData {
	if data == nil {
		return nil
	}
	return &smodel.PersonalData{
		GivenName:  data.GivenName,
		Surname:    data.Surname,
		Patronymic: data.Patronymic,
		City:       data.City,
		Birthdate:  data.Birthdate,
	}
}

func StatsDTOToService(stats *dto.Stats) *smodel.Stats {
	if stats == nil {
		return nil
	}
	return &smodel.Stats{
		UtilizeCount: stats.UtilizeCount,
		ReportCount:  stats.ReportCount,
		Rating:       stats.Rating,
	}
}
