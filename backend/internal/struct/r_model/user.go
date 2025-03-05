package rmodel

import "time"

type User struct {
	Id       uint64
	Creds    *Creds
	Data     *PersonalData
	Stats    *Stats
	Sessions *Session
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

type Session struct {
	SessionId      uint64
	UserId         uint64
	ExpirationTime time.Time
}
