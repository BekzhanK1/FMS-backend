package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"user-service/internal/config"
	"user-service/internal/utils"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteError(w, http.StatusUnauthorized, errors.New("authorization header is missing"))
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Envs.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			utils.WriteError(w, http.StatusUnauthorized, errors.New("invalid or expired token"))
			return
		}

		userID, ok := claims["userID"].(string)
		if !ok {
			utils.WriteError(w, http.StatusUnauthorized, errors.New("invalid token claims"))
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, userID)
		log.Printf("AuthMiddleware: userID: %s", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
