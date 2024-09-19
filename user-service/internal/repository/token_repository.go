package repository

import (
	"database/sql"
	"fmt"
	"user-service/internal/models"
)

func (s *Store) GetTokenByUserId(userId int) (*models.Token, error) {
	query := `SELECT * FROM tokens WHERE userId = $1`

	row := s.db.QueryRow(query, userId)
	
	token := &models.Token{}
	err := row.Scan(
		&token.ID,
		&token.UserID,
		&token.Token,
		&token.Expiration,
		&token.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not get token: %w", err)
	}

	return token, nil
}

func (s *Store) CreateToken(token *models.Token) error {
	query := `
	INSERT INTO tokens (userId, token, expiration, updated_at)
	VALUES ($1, $2, $3, $4)
	`
	_, err := s.db.Exec(query, token.UserID, token.Token, token.Expiration, token.UpdatedAt)

	if err != nil {
		return fmt.Errorf("could not create token: %w", err)
	}

	return nil
}

func (s *Store) UpdateTokenByUserId(userId int, token *models.Token) error {
	query := `
	UPDATE tokens
		SET token = $1, expiration = $2, updated_at = $3
	WHERE userId = $4
`
	_, err := s.db.Exec(query,
		token.Token,
		token.Expiration,
		token.UpdatedAt,
		userId,
	)

	if err != nil {
		return fmt.Errorf("could not update token: %w", err)
	}

	return nil
}
