package smodel

import (
	"time"

	"github.com/Krab1o/meebin/internal/model"
)

type User struct {
	Id       uint64
	Roles    []model.Role
	Creds    *Creds
	Data     *PersonalData
	Stats    *Stats
	Sessions *Tokens
}

type Creds struct {
	Username string
	Email    string
	Password string
}

type PersonalData struct {
	GivenName  string
	Surname    string
	Patronymic string
	City       string
	Birthdate  time.Time
}

type Stats struct {
	UtilizeCount uint64
	ReportCount  uint64
	Rating       float64
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
