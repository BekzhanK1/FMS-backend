package auth

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"user-service/internal/config"
	"user-service/internal/helpers"
	"user-service/internal/middleware"
	"user-service/internal/models"
	"user-service/types"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	tokenStore types.TokenStore
	userStore  types.UserStore
}

func NewService(tokenStore types.TokenStore, userStore types.UserStore) *Service {
	return &Service{
		tokenStore,
		userStore,
	}
}

func (s *Service) Login(email, password string) (types.Tokens, error) {
	user, err := s.userStore.GetUserByEmail(email)
	if err != nil {
		return types.Tokens{}, fmt.Errorf("could not get user by email: %w", err)
	}

	if user == nil {
		return types.Tokens{}, fmt.Errorf("invalid email or password")
	}

	if !user.IsActive {
		return types.Tokens{}, fmt.Errorf("user is not active")
	}

	if err = helpers.CheckPasswordHash(password, user.PasswordHash); err != nil {
		return types.Tokens{}, fmt.Errorf("invalid email or password")
	}

	tokens, err := CreateJWT(user.ID)
	if err != nil {
		return types.Tokens{}, fmt.Errorf("could not create JWT: %w", err)
	}

	token := &models.Token{
		UserID:     user.ID,
		Token:      tokens.RefreshToken,
		Expiration: time.Now().Add(time.Duration(config.Envs.JwtExpRefreshToken) * time.Second),
		UpdatedAt:  time.Now(),
	}

	existingToken, err := s.GetTokenByUserId(user.ID)
	if err != nil {
		return types.Tokens{}, fmt.Errorf("could not get token by user ID: %w", err)
	}

	if existingToken == nil {
		if err := s.CreateToken(token); err != nil {
			return types.Tokens{}, fmt.Errorf("could not create token: %w", err)
		}
	} else {
		if err := s.UpdateTokenByUserId(user.ID, token); err != nil {
			return types.Tokens{}, fmt.Errorf("could not update token: %w", err)
		}
	}

	return tokens, nil

}

func CreateJWT(userID int) (types.Tokens, error) {
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
		return types.Tokens{}, err
	}

	refreshTokenString, err := refreshToken.SignedString(secret)
	if err != nil {
		return types.Tokens{}, err
	}

	return types.Tokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func GetUserIDFromContext(ctx context.Context) (int, error) {
	userIdstr, ok := ctx.Value(middleware.UserKey).(string)
	if !ok {
		return 0, fmt.Errorf("unable to get user ID from context")
	}
	userId, err := strconv.Atoi(userIdstr)
	if err != nil {
		return 0, err
	}

	return userId, nil
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
