package auth

type Token struct {
	AuthorizationToken string
	RefreshToken       string
	ExpiresAt          int64
}

func NewToken(authorizationToken, refreshToken string, expiresAt int64) *Token {
	return &Token{
		AuthorizationToken: authorizationToken,
		RefreshToken:       refreshToken,
		ExpiresAt:          expiresAt,
	}
}
