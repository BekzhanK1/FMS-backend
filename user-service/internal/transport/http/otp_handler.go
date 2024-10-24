package http

import (
	"fmt"
	"net/http"
	"user-service/internal/utils"
)

type ActivateUserRequest struct {
	EncryptedEmail string `json:"key"`
	OTPCode        string `json:"otp_code"`
}

func (h *Handler) ActivateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req ActivateUserRequest
	fmt.Print("Hello")
	err := utils.ParseJSON(r, &req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	encryptedEmail := req.EncryptedEmail
	otpCode := req.OTPCode

	fmt.Printf("ActivateUserHandler: encryptedEmail: %s, otpCode: %s\n", encryptedEmail, otpCode)

	err = h.userService.ActivateUser(encryptedEmail, otpCode)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Account is activated successfully")
}