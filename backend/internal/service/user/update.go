package user

import (
	"context"

	converter "github.com/Krab1o/meebin/internal/converter/service/user"
	rmodel "github.com/Krab1o/meebin/internal/model/user/r_model"
	smodel "github.com/Krab1o/meebin/internal/model/user/s_model"
	"github.com/Krab1o/meebin/internal/service"
	"golang.org/x/crypto/bcrypt"
)

func doUpdate(user *smodel.User) bool {
	return user.Creds != nil && *user.Creds != smodel.Creds{} ||
		user.Data != nil && *user.Data != smodel.PersonalData{}
}

func (s *serv) Update(
	ctx context.Context,
	updaterId uint64,
	user *smodel.User,
) (*smodel.User, error) {
	if updaterId != user.Id {
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

	var updatedRepoUser *rmodel.User
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		repoUser := converter.UserServiceToRepo(user)
		var err error

		if user.Creds != nil {
			err = s.userRepository.UpdateCreds(ctx, user.Id, repoUser.Creds)
			if err != nil {
				return service.ErrorDBToService(err)
			}
		}
		if user.Data != nil {
			err = s.userRepository.UpdatePersonalData(ctx, user.Id, repoUser.Data)
			if err != nil {
				return service.ErrorDBToService(err)
			}
		}
		updatedRepoUser, err = s.userRepository.GetUserById(ctx, user.Id)
		if err != nil {
			return service.ErrorDBToService(err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	updatedUser := converter.UserRepoToService(updatedRepoUser)
	return updatedUser, err
}
