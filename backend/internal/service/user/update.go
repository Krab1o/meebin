package user

import (
	"context"

	convUser "github.com/Krab1o/meebin/internal/converter/service/user"
	smodel "github.com/Krab1o/meebin/internal/model/s_model"
	"github.com/Krab1o/meebin/internal/service"
	"golang.org/x/crypto/bcrypt"
)

func doUpdate(user *smodel.User) bool {
	return user.Creds != nil && *user.Creds != smodel.Creds{} ||
		user.Data != nil && *user.Data != smodel.PersonalData{} ||
		user.Stats != nil && *user.Stats != smodel.Stats{}
}

func (s *serv) Update(ctx context.Context, user *smodel.User) (*smodel.User, error) {
	startUpdate := doUpdate(user)
	if user.Creds != nil && user.Creds != (&smodel.Creds{}) {
		if user.Creds.Password != "" {
			hashedBytes, err := bcrypt.GenerateFromPassword(
				[]byte(user.Creds.Password),
				bcrypt.DefaultCost,
			)
			if err != nil {
				return nil, service.NewInternalError(err)
			}
			user.Creds.Password = string(hashedBytes)
		}
	}
	// log.Println(startUpdate)
	if startUpdate {
		repoUser := convUser.UserServiceToRepo(user)
		var err error
		if user.Creds != nil {
			err = s.userRepo.UpdateCreds(ctx, nil, user.Id, repoUser.Creds)
			if err != nil {
				return nil, service.ErrorDBToService(err, nil)
			}
		}
		if user.Data != nil {
			err = s.userRepo.UpdatePersonalData(ctx, nil, user.Id, repoUser.Data)
			if err != nil {
				return nil, service.ErrorDBToService(err, nil)
			}
		}
		if user.Stats != nil {
			err = s.userRepo.UpdateStats(ctx, nil, user.Id, repoUser.Stats)
			if err != nil {
				return nil, service.ErrorDBToService(err, nil)
			}
		}

		newRepoUser, err := s.userRepo.GetById(ctx, nil, user.Id)
		if err != nil {
			return nil, service.ErrorDBToService(err, nil)
		}
		newUser := convUser.UserRepoToService(newRepoUser)
		return newUser, nil
	} else {
		return nil, service.NewSemanticError(nil, "Unable to process entity")
	}

}
