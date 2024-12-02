package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

const (
	ACCESS_TOKEN_DURATION  = time.Minute * 15
	REFRESH_TOKEN_DURATION = time.Hour * 24 * 7
)

func GenerateToken(username, role string) (*TokenResponse, error) {
	var err error
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"expires":  time.Now().Add(ACCESS_TOKEN_DURATION),
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expires": time.Now().Add(REFRESH_TOKEN_DURATION),
	})

	secretAccessToken := os.Getenv("SECRET_ACCESS_TOKEN")
	secretRefreshToken := os.Getenv("SECRET_REFRESH_TOKEN")

	accessTokenString, err := accessToken.SignedString([]byte(secretAccessToken))
	if err != nil {
		return nil, err
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(secretRefreshToken))
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
