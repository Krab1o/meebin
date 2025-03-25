package update

import (
	"github.com/Krab1o/meebin/internal/model/user/dto"
	smodel "github.com/Krab1o/meebin/internal/model/user/s_model"
)

func UpdatedUserDTOToService(user *dto.UpdateUser) *smodel.User {
	dtoCreds := UpdatedCredsDTOToService(user.Creds)
	dtoData := UpdatedDataDTOToService(user.Data)
	return &smodel.User{
		Id:    user.Id,
		Creds: dtoCreds,
		Data:  dtoData,
	}
}

func UpdatedCredsDTOToService(creds *dto.UpdateCreds) *smodel.Creds {
	if creds == nil {
		return nil
	}
	return &smodel.Creds{
		Username: creds.Username,
		Email:    creds.Email,
		Password: creds.Password,
	}
}

func UpdatedDataDTOToService(data *dto.UpdatePersonalData) *smodel.PersonalData {
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
