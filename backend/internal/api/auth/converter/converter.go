package converter

import (
	"github.com/Krab1o/meebin/internal/model"
	"github.com/Krab1o/meebin/internal/model/dto"
	smodel "github.com/Krab1o/meebin/internal/model/s_model"
)

func RequestUserDTOToService(newUser *dto.RequestCreateUser) *smodel.User {
	creds := CredsDTOToService(newUser.Creds)
	data := DataDTOToService(newUser.Data)
	return &smodel.User{
		Creds: creds,
		Data:  data,
	}
}

func UserDTOToService(user *dto.ResponseProfileUser) *smodel.User {
	creds := CredsDTOToService(user.Creds)
	data := DataDTOToService(user.Data)
	stats := StatsDTOToService(user.Stats)
	roles := RolesDTOToService(user.Roles)
	return &smodel.User{
		Id:    user.Id,
		Roles: roles,
		Creds: creds,
		Data:  data,
		Stats: stats,
	}
}

func RolesDTOToService(roles []model.Role) []model.Role {
	if roles == nil {
		return nil
	}

	copiedRoles := make([]model.Role, len(roles))
	copy(copiedRoles, roles)

	return copiedRoles
}

func CredsDTOToService(creds *dto.Creds) *smodel.Creds {
	return &smodel.Creds{
		Username: creds.Username,
		Email:    creds.Email,
		Password: creds.Password,
	}
}

func DataDTOToService(data *dto.PersonalData) *smodel.PersonalData {
	return &smodel.PersonalData{
		GivenName:  data.GivenName,
		Surname:    data.Surname,
		Patronymic: data.Patronymic,
		City:       data.City,
		Birthdate:  data.Birthdate,
	}
}

func StatsDTOToService(stats *dto.Stats) *smodel.Stats {
	return &smodel.Stats{
		UtilizeCount: stats.UtilizeCount,
		ReportCount:  stats.ReportCount,
		Rating:       stats.Rating,
	}
}

func ResponseTokensServiceToDTO(tokens *smodel.Tokens) *dto.ReponseTokens {
	return &dto.ReponseTokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
}
