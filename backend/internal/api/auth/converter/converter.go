package converter

import (
	"github.com/Krab1o/meebin/internal/api/auth/dto"
	smodel "github.com/Krab1o/meebin/internal/model/s_model"
)

func CredsDTOToService(creds *dto.UserCreds) *smodel.User {
	return &smodel.User{
		Username: creds.Username,
		Password: creds.Password,
	}
}
