package converter

import (
	"github.com/Krab1o/meebin/internal/struct/dto"
	smodel "github.com/Krab1o/meebin/internal/struct/s_model"
)

const (
	defaultID = 0
)

// TODO: possibility of shallow copying. Check for errors and search for better
// solution in case of an error
func UserDTOToService(user *dto.User) *smodel.User {
	creds := CredsDTOToService(user.Creds)
	return &smodel.User{
		Id:    defaultID,
		Creds: creds,
		Data:  (*smodel.PersonalData)(user.Data),
		Stats: (*smodel.Stats)(user.Stats),
	}
}

func CredsDTOToService(creds *dto.Creds) *smodel.Creds {
	return &smodel.Creds{
		Username: creds.Username,
		Email:    creds.Email,
		Password: creds.Password,
	}
}

func TokensServiceToDTO(tokens *smodel.Tokens) *dto.Tokens {
	return (*dto.Tokens)(tokens)
}
