package repository

import (
	"database/sql"
	"fmt"
	"user-service/internal/models"
	"user-service/internal/utils"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db,
	}
}

func (s *UserStore) CreateUser(user *models.User) (*models.User, error) {
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

func (s *UserStore) GetUserById(id int) (*models.User, error) {
	query := `SELECT id, email, username, phone_number, is_active, role, profile_picture_url, created_at, updated_at FROM users WHERE id = $1`

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

func (s *UserStore) UpdateUser(userId int, user *models.User) error {
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

func (s *UserStore) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete user: %w", err)
	}

	return nil
}

func (s *UserStore) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, email, username, password_hash, phone_number, is_active, role, profile_picture_url, created_at, updated_at FROM users WHERE email = $1`

	row := s.db.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Username,
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
