package dto

import (
	"time"

	"github.com/Krab1o/meebin/internal/model"
)

// @Description BaseUser-part of a list-of-users requests
type BaseUser struct {
	Id uint64 `json:"id,omitempty"    example:"42"`
	// Roles' array
	Roles []model.Role      `json:"roles,omitempty"`
	Creds *BaseCreds        `json:"creds"`
	Data  *BasePersonalData `json:"personalData"`
	Stats *BaseStats        `json:"stats"`
}

// @Description BaseUser credentials
type BaseCreds struct {
	Username string `json:"username"           example:"user123"             binding:"omitempty,min=3,max=20"`
	Email    string `json:"email"              example:"user123@example.com" binding:"omitempty,email"                           format:"email"`
	Password string `json:"password,omitempty" example:"Password123"         binding:"omitempty,min=8,digit,uppercase,lowercase" format:"password"`
}

// @Description BaseUser personal data
type BasePersonalData struct {
	GivenName  string    `json:"givenName"  example:"Ivan"`
	Surname    string    `json:"surname"    example:"Ivanov"`
	Patronymic string    `json:"patronymic" example:"Ivanovich"`
	City       string    `json:"city"       example:"Vladivostok"`
	Birthdate  time.Time `json:"birthDate"                        format:"date-time"`
}

// TODO: add validation for stats fields

// @Description BaseUser statistics
type BaseStats struct {
	UtilizeCount uint64  `json:"utilizeCount"`
	ReportCount  uint64  `json:"reportCount"`
	Rating       float64 `json:"rating"`
}
