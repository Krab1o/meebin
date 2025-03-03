package smodel

import "time"

type User struct {
	Id        uint64
	UserCreds Creds
	UserStats Stats
	UserData  PersonalData
}

type Creds struct {
	Username string
	Email    string
	Password string
}

type Stats struct {
	UtilizeCount int64
	ReportCount  int64
	Rating       float64
}

type PersonalData struct {
	GivenName  string
	Surname    string
	Patronymic string
	City       string
	Birthdate  time.Time
}

type Tokens struct {
	RefreshToken string
	AccessToken  string
}

type AccessToken string
