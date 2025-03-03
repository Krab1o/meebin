package auth

import (
	"context"

	smodel "github.com/Krab1o/meebin/internal/struct/s_model"
)

func (s *authService) Refresh(ctx context.Context, id uint64) (*smodel.AccessToken, error) {
	return nil, nil
}
