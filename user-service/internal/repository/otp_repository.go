package repository

import (
	"database/sql"
	"fmt"
	"user-service/internal/models"
	"user-service/internal/utils"
)

func (s *Store) CreateOTP(user *models.User) (string, string, error) {
	encryptedEmail, err := utils.Encrypt(user.Email)
	if err != nil {
		return "", "", fmt.Errorf("could not encrypt email: %w", err)
	}
	otp := &models.OTP{
		UserID:  user.ID,
		OTP_Code: utils.GenerateOTP(),
	}

	query := `
	INSERT INTO otp (user_id, otp_code)
	VALUES ($1, $2)
	`
	_, err = s.db.Exec(query, otp.UserID, otp.OTP_Code)

	if err != nil {
		return "", "", fmt.Errorf("could not create otp: %w", err)
	}

	return otp.OTP_Code, encryptedEmail, nil
}

func (s *Store) DeleteOTP(userId int) error {
	query := `
	DELETE FROM otp
	WHERE user_id = $1
	`
	_, err := s.db.Exec(query, userId)

	if err != nil {
		return fmt.Errorf("could not delete otp: %w", err)
	}

	return nil
}

func (s *Store) GetOTPByUserId(userId int) (*models.OTP, error) {
	query := `
	SELECT user_id, otp_code, expires_at
	FROM otp
	WHERE user_id = $1
	`
	row := s.db.QueryRow(query, userId)

	otp := &models.OTP{}
	err := row.Scan(
		&otp.UserID,
		&otp.OTP_Code,
		&otp.ExpiresAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no OTP found for user ID %d", userId)
		}
		return nil, fmt.Errorf("could not get otp: %w", err)
	}

	return otp, nil
}


func (s *Store) RegenerateOTP(user_id int, otp_code string) error {
	query := `
	UPDATE otp
	SET otp_code = $2, expires_at = CURRENT_TIMESTAMP + INTERVAL '10 minutes'
	WHERE user_id = $1
	`
	_, err := s.db.Exec(query, user_id, otp_code)

	if err != nil {
		return fmt.Errorf("could not regenerate otp: %w", err)
	}

	return nil
}
