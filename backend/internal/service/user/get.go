package user

import (
	"context"

	converter "github.com/Krab1o/meebin/internal/converter/service/user"
	smodel "github.com/Krab1o/meebin/internal/model/user/s_model"
	"github.com/Krab1o/meebin/internal/service"
)

func (s *serv) Get(ctx context.Context, userId uint64) (*smodel.User, error) {
	repoUser, err := s.userRepository.GetUserById(ctx, userId)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}
	user := converter.UserRepoToService(repoUser)
	return user, nil
}
