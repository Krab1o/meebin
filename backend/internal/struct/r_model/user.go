package rmodel

import "time"

type User struct {
	Id       uint64
	Creds    *UserCreds
	Data     *PersonalData
	Stats    *Stats
	Sessions *Tokens
}

type UserCreds struct {
	Username string
	Email    string
	Password string
}

type PersonalData struct {
	FirstName  string
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
	RefreshToken string
	AccessToken  string
}

type AccessToken string
