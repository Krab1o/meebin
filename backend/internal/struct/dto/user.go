package dto

import "time"

type NewUser struct {
	Id    uint64        `json:"id"`
	Creds *Creds        `json:"creds"`
	Data  *PersonalData `json:"personalData"`
}

type User struct {
	Id    uint64        `json:"id"`
	Creds *Creds        `json:"creds"`
	Data  *PersonalData `json:"personalData"`
	Stats *Stats        `json:"stats,omitempty"`
}

type Creds struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PersonalData struct {
	GivenName  string    `json:"givenName"`
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
