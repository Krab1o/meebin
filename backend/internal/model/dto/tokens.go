package dto

type ReponseTokens struct {
	AccessToken  string `json:"accessToken"  binding:"required,jwt"`
	RefreshToken string `json:"refreshToken" binding:"required,jwt"`
}

type RefreshToken struct {
	Refresh string `json:"refreshToken" binding:"required,jwt"`
}

type AccessToken struct {
	Access string `json:"accessToken" binding:"required,jwt"`
}
