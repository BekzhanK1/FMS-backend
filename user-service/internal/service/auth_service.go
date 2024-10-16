package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"user-service/internal/config"
	"user-service/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type contextKey string

const UserKey contextKey = "userID"

func CreateJWT(userID int) (Tokens, error) {
	secret := []byte(config.Envs.JWTSecret)
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

func (s *Service) GetTokenByUserId(userId int) (*models.Token, error) {
    token, err := s.tokenStore.GetTokenByUserId(userId)
    if err != nil {
        return nil, fmt.Errorf("could not retrieve token: %w", err)
    }
    return token, nil
}

func (s *Service) CreateToken(token *models.Token) error {
	if err := s.tokenStore.CreateToken(token); err != nil {
		return fmt.Errorf("could not create token: %w", err)
	}
	return nil
}

func (s *Service) UpdateTokenByUserId(userId int, token *models.Token) error {
	if err := s.tokenStore.UpdateTokenByUserId(userId, token); err != nil {
		return fmt.Errorf("could not update token: %w", err)
	}
	return nil
}