package auth

import (
	"context"

	"github.com/Krab1o/meebin/internal/service"
)

func (s *serv) Logout(ctx context.Context, sessionId uint64) error {
	err := s.sessionRepository.DeleteById(ctx, sessionId)
	if err != nil {
		return service.NewInternalError(err)
	}
	return nil
}
