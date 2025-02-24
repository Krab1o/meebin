package rmodel

import "time"

type User struct {
	ID          int64
	Email       string
	Password    string
	UserStats   Stats
	UserData    PersonalData
	UserSession Session
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
