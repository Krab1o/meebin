package token

import (
	"github.com/Krab1o/meebin/internal/model/dto"
	smodel "github.com/Krab1o/meebin/internal/model/s_model"
)

func TokensServiceToDTO(tokens *smodel.Tokens) *dto.ResponseTokens {
	return &dto.ResponseTokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
}
