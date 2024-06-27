package models

type Token struct {
	AuthorizationToken string `json:"auth_token"`
	RefreshToken       string `json:"refresh_token"`
	ExpiresAt          int64  `json:"expires_at"`
}

type User struct {
	Phone    string
	Email    string
	Password string
}
