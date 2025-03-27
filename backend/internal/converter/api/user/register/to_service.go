package register

import (
	"github.com/Krab1o/meebin/internal/model/user/dto"
	smodel "github.com/Krab1o/meebin/internal/model/user/s_model"
)

func NewUserDTOToService(newUser *dto.RegisterUser) *smodel.User {
	dtoCreds := NewCredsDTOToService(newUser.Creds)
	dtoData := NewDataDTOToService(newUser.Data)
	return &smodel.User{
		Creds: dtoCreds,
		Data:  dtoData,
	}
}

func NewCredsDTOToService(newCreds *dto.RegisterCreds) *smodel.Creds {
	if newCreds == nil {
		return nil
	}
	return &smodel.Creds{
		Username: newCreds.Username,
		Email:    newCreds.Email,
		Password: newCreds.Password,
	}
}

func NewDataDTOToService(newData *dto.RegisterPersonalData) *smodel.PersonalData {
	if newData == nil {
		return nil
	}
	return &smodel.PersonalData{
		GivenName:  newData.GivenName,
		Surname:    newData.Surname,
		Patronymic: newData.Patronymic,
		City:       newData.City,
		Birthdate:  newData.Birthdate,
	}
}
