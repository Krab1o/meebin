package dto

import "time"

type User struct {
	ID       uint64        `json:"id"`
	Username string        `json:"username"`
	Password string        `json:"password"`
	Email    string        `json:"email"`
	Data     PersonalData  `json:"personalData"`
	Stats    PersonalStats `json:"personalStats"`
}
type PersonalData struct {
	Firstname string    `json:"firstname"`
	Surname   string    `json:"surname"`
	Lastname  string    `json:"lastname"`
	BirthDate time.Time `json:"birthDate"`
	City      string    `json:"city"`
}

type PersonalStats struct {
	UtilizeCount int64   `json:"utilizeCount"`
	ReportCount  int64   `json:"reportCount"`
	Rating       float64 `json:"rating"`
}
