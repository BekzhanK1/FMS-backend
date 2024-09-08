package repository

import (
	"database/sql"
	"fmt"
	"user-service/internal/database"
	"user-service/internal/models"
)

func CreateUser(user *models.User) error {
	query := `
		INSERT INTO Users (email, username, phone_number, password_hash, is_active, role, profile_picture_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`
	err := database.DB.QueryRow(query,
		user.Email,
		user.Username,
		user.Phone, // This will map to the phone_number column in the database
		user.PasswordHash,
		user.IsActive,
		user.Role,
		user.ProfilePicture,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func GetUser(id int) (*models.User, error) {
	query := `SELECT id, email, username, phone_number, password_hash, is_active, role, profile_picture_url, created_at, updated_at FROM Users WHERE id = $1`
	row := database.DB.QueryRow(query, id)
	user := &models.User{}
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Phone,
		&user.PasswordHash,
		&user.IsActive,
		&user.Role,
		&user.ProfilePicture,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found
		}
		return nil, fmt.Errorf("could not get user: %w", err)
	}
	return user, nil
}

func UpdateUser(user *models.User) error {
	query := `
		UPDATE Users
		SET email = $1, username = $2, phone = $3, password_hash = $4, is_active = $5, role = $6, profile_picture = $7, updated_at = $8
		WHERE id = $9
	`
	_, err := database.DB.Exec(query,
		user.Email,
		user.Username,
		user.Phone,
		user.PasswordHash,
		user.IsActive,
		user.Role,
		user.ProfilePicture,
		user.UpdatedAt,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("could not update user: %w", err)
	}
	return nil
}

func DeleteUser(id int) error {
	query := `DELETE FROM Users WHERE id = $1`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete user: %w", err)
	}
	return nil
}