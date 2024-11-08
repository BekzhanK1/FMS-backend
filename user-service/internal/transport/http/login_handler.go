package http

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"user-service/internal/config"
	"user-service/internal/helpers"
	"user-service/internal/models"
	"user-service/shared/utils"
	"user-service/types"

	authService "user-service/internal/service/auth"
)

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Printf("LoginHandler: payload: %+v", payload)

	user, err := h.userService.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	if user == nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid email or password"))
		return
	}

	
	if helpers.CheckPasswordHash(payload.Password, user.PasswordHash) != nil{
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid email or password"))
		return
	}
	

	if !user.IsActive{
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("user is not active"))
		return
	}

	tokens, err := authService.CreateJWT(user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	token := &models.Token{
        UserID:     user.ID,
        Token:      tokens.RefreshToken,
        Expiration: time.Now().Add(time.Duration(config.Envs.JwtExpRefreshToken) * time.Second),
        UpdatedAt:  time.Now(),
    }

	existingToken, err := h.authService.GetTokenByUserId(user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if existingToken == nil {
		if err := h.authService.CreateToken(token); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		if err := h.authService.UpdateTokenByUserId(user.ID, token); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
	}

	if err := utils.WriteJSON(w, http.StatusOK, tokens); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

