package dto

import "time"

type User struct {
	Id    uint64        `json:"id"`
	Creds *Creds        `json:"creds"`
	Data  *PersonalData `json:"personalData"`
	Stats *Stats        `json:"stats"`
}

type Creds struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PersonalData struct {
	FirstName  string    `json:"firstName"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic"`
	City       string    `json:"city"`
	Birthdate  time.Time `json:"birthDate"`
}

type Stats struct {
	UtilizeCount uint64  `json:"utilizeCount"`
	ReportCount  uint64  `json:"reportCount"`
	Rating       float64 `json:"rating"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
