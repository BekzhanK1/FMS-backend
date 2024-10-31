package repository

import (
	"database/sql"
	"fmt"
	"user-service/internal/models"
	"user-service/types"
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

func (s *ApplicationStore) ListApplications() ([]*types.ApplicationResponse, error) {
	query := `
    SELECT 
        a.id AS application_id,
        a.status,
        a.rejection_reason,
        a.created_at AS application_created_at,
        
        u.id AS farmer_id,
        u.first_name,
        u.last_name,
        u.username,
        u.email,
        u.phone_number,
        u.profile_picture_url,
        u.role,
        
        fi.rating,
        fi.experience,
        fi.bio,
        
        f.id AS farm_id,
        f.name AS farm_name,
        f.address AS farm_address,
        f.geo_loc AS farm_geo_loc,
        f.size AS farm_size,
        f.crop_types AS farm_crop_types
    FROM applications a
    JOIN users u ON a.farmer_id = u.id
    LEFT JOIN farmer_info fi ON fi.farmer_id = u.id
    JOIN farms f ON a.farm_id = f.id;
    `

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not list applications with details: %w", err)
	}
	defer rows.Close()

	var applications []*types.ApplicationResponse
	for rows.Next() {
		var app types.ApplicationResponse

		// Scan into struct fields
		err = rows.Scan(
			&app.ID,
			&app.Status,
			&app.RejectionReason,
			&app.CreatedAt,
			&app.Farmer.ID,
			&app.Farmer.FirstName,
			&app.Farmer.LastName,
			&app.Farmer.Username,
			&app.Farmer.Email,
			&app.Farmer.Phone,
			&app.Farmer.ProfilePicture,
			&app.Farmer.Role,
			&app.Farmer.Rating,
			&app.Farmer.Experience,
			&app.Farmer.Bio,
			&app.Farm.ID,
			&app.Farm.Name,
			&app.Farm.Address,
			&app.Farm.GeoLoc,
			&app.Farm.Size,
			&app.Farm.CropTypes,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan application details: %w", err)
		}

		applications = append(applications, &app)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error while listing applications with details: %w", err)
	}

	return applications, nil
}

func (s *ApplicationStore) GetApplicationByID(id int) (*types.ApplicationResponse, error) {
	query := `    SELECT 
        a.id AS application_id,
        a.status,
        a.rejection_reason,
        a.created_at AS application_created_at,
        
        u.id AS farmer_id,
        u.first_name,
        u.last_name,
        u.username,
        u.email,
        u.phone_number,
        u.profile_picture_url,
        u.role,
        
        fi.rating,
        fi.experience,
        fi.bio,
        
        f.id AS farm_id,
        f.name AS farm_name,
        f.address AS farm_address,
        f.geo_loc AS farm_geo_loc,
        f.size AS farm_size,
        f.crop_types AS farm_crop_types
    FROM applications a 
    JOIN users u ON a.farmer_id = u.id
    LEFT JOIN farmer_info fi ON fi.farmer_id = u.id
    JOIN farms f ON a.farm_id = f.id
	WHERE a.id = $1;
    `
	app := &types.ApplicationResponse{}
	err := s.db.QueryRow(query, id).Scan(
		&app.ID,
		&app.Status,
		&app.RejectionReason,
		&app.CreatedAt,
		&app.Farmer.ID,
		&app.Farmer.FirstName,
		&app.Farmer.LastName,
		&app.Farmer.Username,
		&app.Farmer.Email,
		&app.Farmer.Phone,
		&app.Farmer.ProfilePicture,
		&app.Farmer.Role,
		&app.Farmer.Rating,
		&app.Farmer.Experience,
		&app.Farmer.Bio,
		&app.Farm.ID,
		&app.Farm.Name,
		&app.Farm.Address,
		&app.Farm.GeoLoc,
		&app.Farm.Size,
		&app.Farm.CropTypes,
	)
	if err != nil {
		return nil, fmt.Errorf("could not get application by ID: %w", err)
	}
	return app, nil
}

func (s *ApplicationStore) ListApplicationsByFarmerID(farmerID int) ([]*types.ApplicationResponse, error) {
	query := `    SELECT 
        a.id AS application_id,
        a.status,
        a.rejection_reason,
        a.created_at AS application_created_at,
        
        u.id AS farmer_id,
        u.first_name,
        u.last_name,
        u.username,
        u.email,
        u.phone_number,
        u.profile_picture_url,
        u.role,
        
        fi.rating,
        fi.experience,
        fi.bio,
        
        f.id AS farm_id,
        f.name AS farm_name,
        f.address AS farm_address,
        f.geo_loc AS farm_geo_loc,
        f.size AS farm_size,
        f.crop_types AS farm_crop_types
    FROM applications a 
    JOIN users u ON a.farmer_id = u.id
    LEFT JOIN farmer_info fi ON fi.farmer_id = u.id
    JOIN farms f ON a.farm_id = f.id
	WHERE u.id = $1;`

	rows, err := s.db.Query(query, farmerID)
	if err != nil {
		return nil, fmt.Errorf("could not get applications by farmer ID: %w", err)
	}
	defer rows.Close()

	var applications []*types.ApplicationResponse
	for rows.Next() {
		var app types.ApplicationResponse

		err = rows.Scan(
			&app.ID,
			&app.Status,
			&app.RejectionReason,
			&app.CreatedAt,
			&app.Farmer.ID,
			&app.Farmer.FirstName,
			&app.Farmer.LastName,
			&app.Farmer.Username,
			&app.Farmer.Email,
			&app.Farmer.Phone,
			&app.Farmer.ProfilePicture,
			&app.Farmer.Role,
			&app.Farmer.Rating,
			&app.Farmer.Experience,
			&app.Farmer.Bio,
			&app.Farm.ID,
			&app.Farm.Name,
			&app.Farm.Address,
			&app.Farm.GeoLoc,
			&app.Farm.Size,
			&app.Farm.CropTypes,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan application details: %w", err)
		}

		applications = append(applications, &app)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error while listing applications with details: %w", err)
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
