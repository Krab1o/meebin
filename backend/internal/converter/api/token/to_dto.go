package token

import (
	"github.com/Krab1o/meebin/internal/model/dto"
	smodel "github.com/Krab1o/meebin/internal/model/s_model"
)

func TokensServiceToDTO(tokens *smodel.Tokens) *dto.ReponseTokens {
	return &dto.ReponseTokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
}
