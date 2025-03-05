package auth

import (
	"context"

	"github.com/Krab1o/meebin/internal/service"
)

func (s *authService) Logout(ctx context.Context, sessionId uint64) error {
	err := s.sessionRepo.DeleteSession(ctx, nil, sessionId)
	if err != nil {
		return service.NewInternalError(err)
	}
	return nil
}
