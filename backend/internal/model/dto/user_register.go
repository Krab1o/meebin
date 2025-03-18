package dto

import "time"

// @Description New registrating user structure
type NewUser struct {
	Creds *NewCreds        `json:"creds"        binding:"required"`
	Data  *NewPersonalData `json:"personalData" binding:"required"`
}

// @Description New user's credentials
type NewCreds struct {
	Username string `json:"username"           example:"user123"             binding:"required,min=3,max=20"`
	Email    string `json:"email"              example:"user123@example.com" binding:"required,email"                           format:"email"`
	Password string `json:"password,omitempty" example:"Password123"         binding:"required,min=8,digit,uppercase,lowercase" format:"password"`
}

// @Description New user's personal data
type NewPersonalData struct {
	GivenName  string `json:"givenName"  example:"Ivan"        binding:"required"`
	Surname    string `json:"surname"    example:"Ivanov"      binding:"required"`
	Patronymic string `json:"patronymic" example:"Ivanovich"   binding:"required"`
	City       string `json:"city"       example:"Vladivostok" binding:"required"`
	// This field accepts default time format from RFC (DD-MM-YYYY, just pick default time library)
	Birthdate time.Time `json:"birthDate"                        binding:"required" format:"date-time"`
}
