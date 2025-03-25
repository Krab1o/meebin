package update

import (
	"github.com/Krab1o/meebin/internal/model/user/dto"
	smodel "github.com/Krab1o/meebin/internal/model/user/s_model"
)

func UpdatedUserServiceToDTO(user *smodel.User) *dto.UpdateUser {
	dtoCreds := UpdatedCredsServiceToDTO(user.Creds)
	dtoData := UpdatedDataServiceToDTO(user.Data)
	return &dto.UpdateUser{
		Id:    user.Id,
		Creds: dtoCreds,
		Data:  dtoData,
	}
}

func UpdatedCredsServiceToDTO(creds *smodel.Creds) *dto.UpdateCreds {
	if creds == nil {
		return nil
	}
	return &dto.UpdateCreds{
		Username: creds.Username,
		Email:    creds.Email,
	}
}

func UpdatedDataServiceToDTO(data *smodel.PersonalData) *dto.UpdatePersonalData {
	if data == nil {
		return nil
	}
	return &dto.UpdatePersonalData{
		GivenName:  data.GivenName,
		Surname:    data.Surname,
		Patronymic: data.Patronymic,
		Birthdate:  data.Birthdate,
		City:       data.City,
	}
}
