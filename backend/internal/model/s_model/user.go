package smodel

import "time"

type User struct {
	Id           int64
	Email        string
	Username     string
	Password     string
	PersonalData struct {
		FirstName string
		SurName   string
		LastName  string
		City      string
		Birthdate time.Time
	}
	Stats struct {
		UtilizeCount int64
		ReportCount  int64
		Rating       float64
	}
}

type UserCreds struct {
	Username string
	Password string
}

type Tokens struct {
	RefreshToken string
	AccessToken  string
}
