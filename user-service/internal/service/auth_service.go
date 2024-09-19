package service

import (
	"context"
	"strconv"
	"time"
	"user-service/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type contextKey string

const UserKey contextKey = "userID"

func CreateJWT(secret []byte, userID int) (Tokens, error) {
	accessTokenExp := config.Envs.JwtExpAccessToken
	access_token_expiration := time.Second * time.Duration(accessTokenExp)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(access_token_expiration).Unix(),
	})

	refreshTokenExp := config.Envs.JwtExpRefreshToken
	refreshTokenExpiration := time.Second * time.Duration(refreshTokenExp)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(refreshTokenExpiration).Unix(),
	})

	accessTokenString, err := accessToken.SignedString(secret)
	if err != nil {
		return Tokens{}, err
	}

	refreshTokenString, err := refreshToken.SignedString(secret)
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{accessTokenString, refreshTokenString}, nil
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userID
}
