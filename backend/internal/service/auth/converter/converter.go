package converter

import (
	rmodel "github.com/Krab1o/meebin/internal/struct/r_model"
	smodel "github.com/Krab1o/meebin/internal/struct/s_model"
	"golang.org/x/crypto/bcrypt"
)

func UserServiceToRepository(user *smodel.User) (*rmodel.User, error) {
	repoCreds, err := CredsServiceToRepo(user.Creds)
	if err != nil {
		return nil, err
	}
	return &rmodel.User{
		Id:    user.Id,
		Creds: repoCreds,
		Stats: (*rmodel.Stats)(user.Stats),
		Data:  (*rmodel.PersonalData)(user.Data),
	}, nil
}

func CredsServiceToRepo(creds *smodel.Creds) (*rmodel.Creds, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(creds.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	return &rmodel.Creds{
		Username: creds.Username,
		Email:    creds.Email,
		Password: string(hashedBytes),
	}, nil
}

// func
