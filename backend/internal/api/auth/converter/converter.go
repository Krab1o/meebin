package converter

import (
	"github.com/Krab1o/meebin/internal/struct/dto"
	smodel "github.com/Krab1o/meebin/internal/struct/s_model"
)

const (
	startUtilizeCount = 0
	startReportCount  = 0
	startRating       = 0.0
)

// TODO: possibility of shallow copying. Check for errors and search for better
// solution in case of an error

func NewUserDTOToService(newUser *dto.NewUser) *smodel.User {
	creds := CredsDTOToService(newUser.Creds)
	data := DataDTOToService(newUser.Data)
	return &smodel.User{
		Creds: creds,
		Data:  data,
	}
}

func UserDTOToService(user *dto.User) *smodel.User {
	creds := CredsDTOToService(user.Creds)
	data := DataDTOToService(user.Data)
	stats := StatsDTOToService(user.Stats)
	return &smodel.User{
		Id:    user.Id,
		Creds: creds,
		Data:  data,
		Stats: stats,
	}
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

func TokensServiceToDTO(tokens *smodel.Tokens) *dto.Tokens {
	return (*dto.Tokens)(tokens)
}
