package token

import (
	"github.com/Krab1o/meebin/internal/model/user/dto"
	smodel "github.com/Krab1o/meebin/internal/model/user/s_model"
)

func TokensServiceToDTO(tokens *smodel.Tokens) *dto.ResponseTokens {
	return &dto.ResponseTokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
}
