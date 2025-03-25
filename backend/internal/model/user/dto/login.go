package dto

type LoginCreds struct {
	Email    string `json:"email"              example:"user123@example.com" binding:"required,email"                           format:"email"`
	Password string `json:"password,omitempty" example:"Password123"         binding:"required,min=8,digit,uppercase,lowercase" format:"password"`
}
