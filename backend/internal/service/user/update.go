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
		user.Data != nil && *user.Data != smodel.PersonalData{}
}

// TODO: construct map with repository column name
func (s *serv) Update(
	ctx context.Context,
	user *smodel.User,
	updatedUserId uint64,
) (*smodel.User, error) {
	if user.Id != updatedUserId {
		return nil, service.NewForbiddenError(nil)
	}
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

	if !startUpdate {
		return nil, service.NewNoUpdateError(nil)
	}

	repoUser := convUser.UserServiceToRepo(user)
	var err error
	if user.Creds != nil {
		err = s.userRepo.UpdateCreds(ctx, nil, user.Id, repoUser.Creds)
		if err != nil {
			return nil, service.ErrorDBToService(err)
		}
	}
	if user.Data != nil {
		err = s.userRepo.UpdatePersonalData(ctx, nil, user.Id, repoUser.Data)
		if err != nil {
			return nil, service.ErrorDBToService(err)
		}
	}

	newRepoUser, err := s.userRepo.GetById(ctx, nil, user.Id)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}
	newUser := convUser.UserRepoToService(newRepoUser)
	return newUser, nil
}
