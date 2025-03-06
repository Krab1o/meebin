package dto

import "time"

type RequestCreateUser struct {
	Creds *Creds        `json:"creds"        binding:"required"`
	Data  *PersonalData `json:"personalData" binding:"required"`
}

type ResponseProfileUser struct {
	Id    uint64        `json:"id"`
	Role  string        `json:"role"`
	Creds *Creds        `json:"creds"`
	Data  *PersonalData `json:"personalData"`
	Stats *Stats        `json:"stats,omitempty"`
}

type Creds struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,digit,uppercase,lowercase"`
}

type PersonalData struct {
	GivenName  string    `json:"givenName"  binding:"required"`
	Surname    string    `json:"surname"    binding:"required"`
	Patronymic string    `json:"patronymic"`
	City       string    `json:"city"       binding:"required"`
	Birthdate  time.Time `json:"birthDate"  binding:"required"`
}

type Stats struct {
	UtilizeCount uint64  `json:"utilizeCount"`
	ReportCount  uint64  `json:"reportCount"`
	Rating       float64 `json:"rating"`
}
