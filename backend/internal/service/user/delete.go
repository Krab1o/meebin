package user

import (
	"context"

	"github.com/Krab1o/meebin/internal/service"
)

func (s *serv) Delete(ctx context.Context, userId uint64) error {
	err := s.userRepository.DeleteById(ctx, userId)
	if err != nil {
		return service.ErrorDBToService(err)
	}
	return nil
}
