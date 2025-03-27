package login

import (
	"github.com/Krab1o/meebin/internal/model/user/dto"
	smodel "github.com/Krab1o/meebin/internal/model/user/s_model"
)

func LoginCredsDTOToService(creds *dto.LoginCreds) *smodel.Creds {
	return &smodel.Creds{
		Email:    creds.Email,
		Password: creds.Password,
	}
}
