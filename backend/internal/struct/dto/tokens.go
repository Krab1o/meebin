package dto

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Token struct {
	Refresh string `json:"refreshToken"`
}
