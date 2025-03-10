package auth

import (
	"context"

	smodel "github.com/Krab1o/meebin/internal/model/s_model"
)

func (s *authService) Profile(ctx context.Context, userId uint64) (*smodel.User, error) {
	s.userRepo.GetUserById(userId)
}
