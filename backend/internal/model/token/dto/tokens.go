package dto

type ResponseTokens struct {
	// Should be a JWT access-token
	AccessToken string `json:"accessToken"  binding:"required,jwt"`
	// Should be a JWT refresh-token
	RefreshToken string `json:"refreshToken" binding:"required,jwt"`
}

type RefreshToken struct {
	// Should be a JWT refresh-token
	Refresh string `json:"refreshToken" binding:"required,jwt"`
}

type AccessToken struct {
	// Should be a JWT access-token
	Access string `json:"accessToken" binding:"required,jwt"`
}
