package rmodel

import "time"

type User struct {
	ID          uint64
	UserCreds   Creds
	UserStats   Stats
	UserData    PersonalData
	UserSession Session
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
	Partonymic string
	City       string
	Birthdate  time.Time
}

type Session struct {
	ID             uint64
	UserID         uint64
	refreshToken   string
	expirationTime time.Time
}
