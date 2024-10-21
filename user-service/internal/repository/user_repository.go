package repository

import (
	"database/sql"
	"fmt"
	"user-service/internal/models"
	"user-service/internal/utils"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db,
	}
}

func (s *Store) CreateUser(user *models.User) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(user.PasswordHash)
	if err != nil {
		return &models.User{}, fmt.Errorf("could not hash password: %w", err)
	}

	query := `
		INSERT INTO users (email, username, first_name, last_name, phone_number, password_hash, is_active, role, profile_picture_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`
	var userID int
	err = s.db.QueryRow(query,
		user.Email,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Phone,
		hashedPassword,
		user.IsActive,
		user.Role,
		user.ProfilePicture,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&userID)

	if err != nil {
		return &models.User{}, fmt.Errorf("could not create user: %w", err)
	}

	user.ID = userID
	return user, nil
}



func (s *Store) GetUserById(id int) (*models.User, error) {
	query := `SELECT id, email, username, first_name, last_name, phone_number, is_active, role, profile_picture_url, created_at, updated_at FROM users WHERE id = $1`

	row := s.db.QueryRow(query, id)
	
	user := &models.User{}
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.IsActive,
		&user.Role,
		&user.ProfilePicture,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not get user: %w", err)
	}

	return user, nil
}

func (s *Store) UpdateUser(userId int, user *models.User) error {
	query := `
		UPDATE users
		SET email = $1, username = $2, first_name = $3, last_name = $4, phone = $5, password_hash = $6, is_active = $7, role = $8, profile_picture = $9, updated_at = $10
		WHERE id = $11
	`
	_, err := s.db.Exec(query,
		user.Email,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Phone,
		user.PasswordHash,
		user.IsActive,
		user.Role,
		user.ProfilePicture,
		user.UpdatedAt,
		userId,
	)

	if err != nil {
		return fmt.Errorf("could not update user: %w", err)
	}

	return nil
}


func (s *Store) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete user: %w", err)
	}

	return nil
}

func (s *Store) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, email, username, first_name, last_name, password_hash, phone_number, is_active, role, profile_picture_url, created_at, updated_at FROM users WHERE email = $1`

	row := s.db.QueryRow(query, email)
	
	user := &models.User{}
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.PasswordHash,
		&user.Phone,
		&user.IsActive,
		&user.Role,
		&user.ProfilePicture,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not get user: %w", err)
	}

	return user, nil
}


func (s *Store) ActivateUser(encryptedEmail string, otpCode string) error {
	userEmail, err := utils.Decrypt(encryptedEmail)
	if err != nil {
		return fmt.Errorf("could not decrypt email: %w", err)
	}

	user, err := s.GetUserByEmail(userEmail)

	if user.IsActive{
		return fmt.Errorf("user is already activated: %w", err)
	}

	otp, err := s.GetOTPByUserId(user.ID)
	if err != nil {
		return fmt.Errorf("could not get token: %w", err)
	}

	if otp.ExpiresAt.Before(utils.GetCurrentTime()) {
		newOtpCode := utils.GenerateOTP()
		err = s.RegenerateOTP(user.ID, newOtpCode)
		if err != nil {
			return fmt.Errorf("couldn't regenerate OTP: %w", err)
		}

		utils.SendEmail(user.Email, "OTP Code", newOtpCode)

		return fmt.Errorf("OTP is expired, a new OTP has been sent to email %s", user.Email)
	}

	fmt.Printf("otpCode: %s, otp.OTP_Code: %s\n", otpCode, otp.OTP_Code)

	if otpCode != otp.OTP_Code {
		return fmt.Errorf("invalid OTP code")
	}

	query := `UPDATE users SET is_active = true WHERE id = $1`
	_, err = s.db.Exec(query, user.ID)
	if err != nil {
		return fmt.Errorf("could not activate user: %w", err)
	}

	err = s.DeleteOTP(user.ID)
	if err != nil {
		return fmt.Errorf("could not delete OTP: %w", err)
	}

	return nil

}

