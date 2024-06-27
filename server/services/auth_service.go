package services

import (
	"context"
	"server/internal/models/auth"
	"server/internal/models/user"
	"server/pkg/jwt"
	"time"

	jwt2 "github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	userOps                *user.Ops
	secret                 []byte
	tokenExpiration        uint
	refreshTokenExpiration uint
}

func NewAuthService(userOps *user.Ops, secret []byte,
	tokenExpiration uint, refreshTokenExpiration uint) *AuthService {
	return &AuthService{
		userOps:                userOps,
		secret:                 secret,
		tokenExpiration:        tokenExpiration,
		refreshTokenExpiration: refreshTokenExpiration,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, user *user.User) (*user.User, *auth.Token, error) {
	createdUser, err := s.userOps.Create(ctx, user)
	if err != nil {
		return nil, nil, err
	}
	var (
		authExp    = time.Now().Add(time.Minute * time.Duration(s.tokenExpiration))
		refreshExp = time.Now().Add(time.Minute * time.Duration(s.refreshTokenExpiration))
	)
	authToken, err := jwt.CreateToken(s.secret, s.userClaims(createdUser, authExp))
	if err != nil {
		return nil, nil, err
	}
	
  refreshToken, err := jwt.CreateToken(s.secret, s.userClaims(createdUser, refreshExp))
	if err != nil {
		return nil, nil, err
	}
	token := auth.NewToken(authToken, refreshToken, authExp.Unix())

	return createdUser, token, nil
}

func (s *AuthService) LoginUser(ctx context.Context, email, pass string) (*auth.Token, error) {
	loggedInUser, err := s.userOps.GetUser(ctx, email, pass)
	if err != nil {
		return nil, err
	}

	// calc expiration time values
	var (
		authExp    = time.Now().Add(time.Minute * time.Duration(s.tokenExpiration))
		refreshExp = time.Now().Add(time.Minute * time.Duration(s.refreshTokenExpiration))
	)

	authToken, err := jwt.CreateToken(s.secret, s.userClaims(loggedInUser, authExp))
	if err != nil {
		return nil, err // todo
	}

	refreshToken, err := jwt.CreateToken(s.secret, s.userClaims(loggedInUser, refreshExp))
	if err != nil {
		return nil, err // todo
	}

	return auth.NewToken(authToken, refreshToken, authExp.Unix()), nil
}

func (s *AuthService) userClaims(user *user.User, exp time.Time) *jwt.UserClaims {
	return &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: &jwt2.NumericDate{
				Time: exp,
			},
		},
		UserID:  user.ID,
		IsAdmin: user.IsAdmin,
	}
}
