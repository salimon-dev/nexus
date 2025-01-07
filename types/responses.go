package types

type AuthResponse struct {
	AccessToken  string     `json:"accessToken"`
	RefreshToken string     `json:"refreshToken"`
	ExpiresAt    string     `json:"expiresAt"`
	Data         PublicUser `json:"data"`
}
