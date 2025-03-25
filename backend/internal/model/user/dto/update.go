package dto

import "time"

// @Description Entity which will be updated
type UpdateUser struct {
	Id    uint64              `json:"id,omitempty" example:"42"`
	Creds *UpdateCreds        `json:"creds"`
	Data  *UpdatePersonalData `json:"personalData"`
}

// @Description Credentials which will be updated
type UpdateCreds struct {
	Username string `json:"username"           example:"user123"             binding:"omitempty,min=3,max=20"`
	Email    string `json:"email"              example:"user123@example.com" binding:"omitempty,email"                           format:"email"`
	Password string `json:"password,omitempty" example:"Password123"         binding:"omitempty,min=8,digit,uppercase,lowercase" format:"password"`
}

// @Description Personal data which will be updated
type UpdatePersonalData struct {
	GivenName  string `json:"givenName"  example:"Ivan"`
	Surname    string `json:"surname"    example:"Ivanov"`
	Patronymic string `json:"patronymic" example:"Ivanovich"`
	City       string `json:"city"       example:"Vladivostok"`
	// This field accepts default time format from RFC (DD-MM-YYYY, just pick default time library)
	Birthdate time.Time `json:"birthDate"                        format:"date"`
}
