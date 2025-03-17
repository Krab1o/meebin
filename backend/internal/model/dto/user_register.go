package dto

import "time"

type NewUser struct {
	Creds *NewCreds        `json:"creds"        binding:"required"`
	Data  *NewPersonalData `json:"personalData" binding:"required"`
}

type NewCreds struct {
	Username string `json:"username"           binding:"required,min=3,max=20"`
	Email    string `json:"email"              binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,min=8,digit,uppercase,lowercase"`
}

type NewPersonalData struct {
	GivenName  string    `json:"givenName"  binding:"required"`
	Surname    string    `json:"surname"    binding:"required"`
	Patronymic string    `json:"patronymic" binding:"required"`
	City       string    `json:"city"       binding:"required"`
	Birthdate  time.Time `json:"birthDate"  binding:"required"`
}
