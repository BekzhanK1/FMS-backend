package handler

import (
	"fmt"
	"net/http"
	"time"
	"user-service/internal/config"
	"user-service/internal/models"
	"user-service/internal/service"
	"user-service/internal/utils"
	"user-service/types"
)

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err) // User not found or another error
		return
	}

	if utils.CheckPasswordHash(payload.Password, user.PasswordHash) != nil{
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid email or password"))
		return
	}

	tokens, err := service.CreateJWT(user.ID)
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

	existingToken, err := h.service.GetTokenByUserId(user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if existingToken == nil {
		if err := h.service.CreateToken(token); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		if err := h.service.UpdateTokenByUserId(user.ID, token); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
	}

	if err := utils.WriteJSON(w, http.StatusOK, tokens); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}