package user

import (
	"context"

	converter "github.com/Krab1o/meebin/internal/converter/service/user"
	smodel "github.com/Krab1o/meebin/internal/model/s_model"
	"github.com/Krab1o/meebin/internal/service"
)

func (s *serv) GetUser(ctx context.Context, userId uint64) (*smodel.User, error) {
	repoUser, err := s.userRepo.GetById(ctx, nil, userId)
	if err != nil {
		return nil, service.ErrorDBToService(err, nil)
	}
	user := converter.UserRepoToService(repoUser)
	return user, nil
}
