package auth

type Token struct {
	AuthorizationToken string `json:"auth_token"`
	RefreshToken       string `json:"refresh_token"`
	ExpiresAt          int64  `json:"expires_at"`
}

func NewToken(authorizationToken, refreshToken string, expiresAt int64) *Token {
	return &Token{
		AuthorizationToken: authorizationToken,
		RefreshToken:       refreshToken,
		ExpiresAt:          expiresAt,
	}
}
