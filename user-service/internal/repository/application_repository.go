package repository

import (
	"database/sql"
	"fmt"
	"user-service/internal/models"
)

type ApplicationStore struct {
	db *sql.DB
}

func NewApplicationStore(db *sql.DB) *ApplicationStore {
	return &ApplicationStore{
		db: db,
	}
}

func (s *ApplicationStore) CreateApplication(application *models.Application) error {
	application.Status = "pending"
	query := `INSERT INTO applications (farmer_id, farm_id, status) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := s.db.QueryRow(query, application.FarmerID, application.FarmID, application.Status).Scan(&application.ID, &application.CreatedAt)
	if err != nil {
		return fmt.Errorf("could not create application: %w", err)
	}

	fmt.Println("Application created successfully")
	return nil
}

func (s *ApplicationStore) ListApplications() ([]*models.Application, error) {
	query := `SELECT id, farmer_id, farm_id, status, rejection_reason, created_at FROM applications`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not list applications: %w", err)
	}
	defer rows.Close()

	var applications []*models.Application
	for rows.Next() {
		application := &models.Application{}
		err = rows.Scan(
			&application.ID,
			&application.FarmerID,
			&application.FarmID,
			&application.Status,
			&application.RejectionReason,
			&application.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan application: %w", err)
		}
		applications = append(applications, application)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error while listing applications: %w", err)
	}

	return applications, nil
}

func (s *ApplicationStore) UpdateApplication(id int, status string, rejectionReason *string) error {
	query := `UPDATE applications SET status = $1, rejection_reason = $2 WHERE id = $3`
	_, err := s.db.Exec(query, status, rejectionReason, id)
	if err != nil {
		return fmt.Errorf("could not update application: %w", err)
	}
	fmt.Println("Application updated successfully")
	return nil
}

func (s *ApplicationStore) GetApplicationByID(id int) (*models.Application, error) {
	query := `SELECT id, farmer_id, farm_id, status, rejection_reason, created_at FROM applications WHERE id = $1`
	application := &models.Application{}
	err := s.db.QueryRow(query, id).Scan(
		&application.ID,
		&application.FarmerID,
		&application.FarmID,
		&application.Status,
		&application.RejectionReason,
		&application.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil 
		}
		return nil, fmt.Errorf("could not get application by ID: %w", err)
	}

	return application, nil
}

func (s *ApplicationStore) GetApplicationByFarmID(farmID int) (*models.Application, error) {
	query := `SELECT id, farmer_id, farm_id, status, rejection_reason, created_at FROM applications WHERE farm_id = $1`
	application := &models.Application{}
	err := s.db.QueryRow(query, farmID).Scan(
		&application.ID,
		&application.FarmerID,
		&application.FarmID,
		&application.Status,
		&application.RejectionReason,
		&application.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not get application by farm ID: %w", err)
	}

	return application, nil
}

func (s *ApplicationStore) GetApplicationsByFarmerID(farmerID int) ([]*models.Application, error) {
	query := `SELECT id, farmer_id, farm_id, status, rejection_reason, created_at FROM applications WHERE farmer_id = $1`
	rows, err := s.db.Query(query, farmerID)
	if err != nil {
		return nil, fmt.Errorf("could not get applications by farmer ID: %w", err)
	}
	defer rows.Close()

	var applications []*models.Application
	for rows.Next() {
		application := &models.Application{}
		err = rows.Scan(
			&application.ID,
			&application.FarmerID,
			&application.FarmID,
			&application.Status,
			&application.RejectionReason,
			&application.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan application: %w", err)
		}
		applications = append(applications, application)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error while getting applications by farmer ID: %w", err)
	}

	return applications, nil
}
