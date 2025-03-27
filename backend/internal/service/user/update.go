package user

import (
	"context"

	converter "github.com/Krab1o/meebin/internal/converter/service/user"
	rmodel "github.com/Krab1o/meebin/internal/model/user/r_model"
	smodel "github.com/Krab1o/meebin/internal/model/user/s_model"
	"github.com/Krab1o/meebin/internal/service"
	"golang.org/x/crypto/bcrypt"
)

func doCredsUpdate(user *smodel.User) bool {
	return user.Creds != nil && *user.Creds != smodel.Creds{}
}

func doDataUpdate(user *smodel.User) bool {
	return user.Data != nil && *user.Data != smodel.PersonalData{}
}

func (s *serv) Update(
	ctx context.Context,
	updaterId uint64,
	user *smodel.User,
) (*smodel.User, error) {
	if updaterId != user.Id {
		return nil, service.NewForbiddenError(nil)
	}

	credsUpdate := doCredsUpdate(user)
	dataUpdate := doDataUpdate(user)
	if !credsUpdate && !dataUpdate {
		return nil, service.NewNoUpdateError(nil)
	}

	if credsUpdate {
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

	var updatedRepoUser *rmodel.User
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		repoUser := converter.UserServiceToRepo(user)
		var err error

		if credsUpdate {
			err = s.userRepository.UpdateCreds(ctx, user.Id, repoUser.Creds)
			if err != nil {
				return service.ErrorDBToService(err)
			}
		}
		if dataUpdate {
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
