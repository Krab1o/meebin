package user

import (
	"context"

	converter "github.com/Krab1o/meebin/internal/converter/service/user"
	smodel "github.com/Krab1o/meebin/internal/model/s_model"
	"github.com/Krab1o/meebin/internal/service"
)

func (s *serv) ListUser(ctx context.Context) ([]smodel.User, error) {
	repoUsers, err := s.userRepo.List(ctx, nil)
	if err != nil {
		return nil, service.ErrorDBToService(err, nil)
	}
	users := make([]smodel.User, len(repoUsers))
	for i, repoUser := range repoUsers {
		users[i] = *converter.UserRepoToService(&repoUser)
	}
	return users, nil
}
