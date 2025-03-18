package updateduser

import (
	"github.com/Krab1o/meebin/internal/model/dto"
	smodel "github.com/Krab1o/meebin/internal/model/s_model"
)

func UpdatedUserServiceToDTO(user *smodel.User) *dto.UpdatedUser {
	dtoCreds := UpdatedCredsServiceToDTO(user.Creds)
	dtoData := UpdatedDataServiceToDTO(user.Data)
	return &dto.UpdatedUser{
		Id:    user.Id,
		Creds: dtoCreds,
		Data:  dtoData,
	}
}

func UpdatedCredsServiceToDTO(creds *smodel.Creds) *dto.UpdatedCreds {
	if creds == nil {
		return nil
	}
	return &dto.UpdatedCreds{
		Username: creds.Username,
		Email:    creds.Email,
	}
}

func UpdatedDataServiceToDTO(data *smodel.PersonalData) *dto.UpdatedPersonalData {
	if data == nil {
		return nil
	}
	return &dto.UpdatedPersonalData{
		GivenName:  data.GivenName,
		Surname:    data.Surname,
		Patronymic: data.Patronymic,
		Birthdate:  data.Birthdate,
		City:       data.City,
	}
}
