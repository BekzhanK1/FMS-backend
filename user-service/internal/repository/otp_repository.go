package repository

import (
	"fmt"
	"user-service/internal/models"
	"user-service/internal/utils"
)

func (s *Store) CreateOTP(userId int) (string, error) {
	otp := &models.OTP{
		UserID:  userId,
		OTP_Code: utils.GenerateOTP(),
	}
	
	query := `
	INSERT INTO otp (user_id, otp_code)
	VALUES ($1, $2)
	`
	_, err := s.db.Exec(query, otp.UserID, otp.OTP_Code)

	if err != nil {
		return "", fmt.Errorf("could not create otp: %w", err)
	}

	return otp.OTP_Code, nil
}

func (s *Store) DeleteOTP(user_id int) error {
	query := `
	DELETE FROM otp
	WHERE user_id = $1
	`
	_, err := s.db.Exec(query, user_id)

	if err != nil {
		return fmt.Errorf("could not delete otp: %w", err)
	}

	return nil
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
