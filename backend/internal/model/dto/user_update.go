package dto

import (
	"time"

	"github.com/Krab1o/meebin/internal/model"
)

type User struct {
	Id    uint64        `json:"id,omitempty"`
	Roles []model.Role  `json:"roles,omitempty"`
	Creds *Creds        `json:"creds"`
	Data  *PersonalData `json:"personalData"`
	Stats *Stats        `json:"stats"`
}

type Creds struct {
	Username string `json:"username"           binding:"omitempty,min=3,max=20"`
	Email    string `json:"email"              binding:"omitempty,email"`
	Password string `json:"password,omitempty" binding:"omitempty,min=8,digit,uppercase,lowercase"`
}

type PersonalData struct {
	GivenName  string    `json:"givenName"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic"`
	City       string    `json:"city"`
	Birthdate  time.Time `json:"birthDate"`
}

// TODO: add validation for fields
type Stats struct {
	UtilizeCount uint64  `json:"utilizeCount"`
	ReportCount  uint64  `json:"reportCount"`
	Rating       float64 `json:"rating"`
}
