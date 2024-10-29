package repository

import (
	"database/sql"
	"fmt"
	"user-service/internal/models"
)

type FarmStore struct {
	db *sql.DB
}

func NewFarmStore(db *sql.DB) *FarmStore {
	return &FarmStore{
		db: db,
	}
}

func (s *FarmStore) CreateFarm(farm *models.Farm) error {
	query := `INSERT INTO farms (name, address, farmer_id, geo_loc, size, crop_types, is_verified, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := s.db.Exec(query, farm.Name, farm.Address, farm.FarmerID, farm.GeoLoc, farm.Size, farm.CropTypes, farm.IsVerified, farm.CreatedAt, farm.UpdatedAt)

	if err != nil {
		return fmt.Errorf("could not create farm: %w", err)
	}

	fmt.Println("Farm created successfully")
	return nil
}

func (s *FarmStore) GetFarmByID(id int) (*models.Farm, error) {
	query := `SELECT id, farmer_id, name, address, geo_loc, size, crop_types, is_verified, created_at, updated_at FROM farms WHERE id = $1`
	row := s.db.QueryRow(query, id)

	farm := &models.Farm{}
	err := row.Scan(
		&farm.ID,
		&farm.FarmerID,
		&farm.Name,
		&farm.Address,
		&farm.GeoLoc,
		&farm.Size,
		&farm.CropTypes,
		&farm.IsVerified,
		&farm.CreatedAt,
		&farm.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not get farm: %w", err)
	}

	return farm, nil
}

func (s *FarmStore) UpdateFarm(farmerID int, farm *models.Farm) error {
	query := `UPDATE farms SET name = $1, address = $2, geo_loc = $3, size = $4, crop_types = $5, is_verified = $6, updated_at = $7 WHERE id = $8 AND farmer_id = $9`
	_, err := s.db.Exec(query, farm.Name, farm.Address, farm.GeoLoc, farm.Size, farm.CropTypes, farm.IsVerified, farm.UpdatedAt, farm.ID, farmerID)

	if err != nil {
		return fmt.Errorf("could not update farm: %w", err)
	}

	fmt.Println("Farm updated successfully")
	return nil
}

func (s *FarmStore) DeleteFarm(farmerID int, id int) error {
	query := `DELETE FROM farms WHERE id = $1 AND farmer_id = $2`
	_, err := s.db.Exec(query, id, farmerID)

	if err != nil {
		return fmt.Errorf("could not delete farm: %w", err)
	}

	fmt.Println("Farm deleted successfully")
	return nil
}

func (s *FarmStore) ListFarms() ([]*models.Farm, error) {
	query := `SELECT id, farmer_id, name, address, geo_loc, size, crop_types, is_verified, created_at, updated_at FROM farms`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not list farms: %w", err)
	}
	defer rows.Close()

	var farms []*models.Farm
	for rows.Next() {
		farm := &models.Farm{}
		err = rows.Scan(
			&farm.ID,
			&farm.FarmerID,
			&farm.Name,
			&farm.Address,
			&farm.GeoLoc,
			&farm.Size,
			&farm.CropTypes,
			&farm.IsVerified,
			&farm.CreatedAt,
			&farm.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan farm: %w", err)
		}
		farms = append(farms, farm)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return farms, nil
}

func (s *FarmStore) ListFarmsByFarmerID(farmerID int) ([]*models.Farm, error) {
	query := `SELECT id, farmer_id, name, address, geo_loc, size, crop_types, is_verified, created_at, updated_at FROM farms WHERE farmer_id = $1`
	rows, err := s.db.Query(query, farmerID)
	if err != nil {
		return nil, fmt.Errorf("could not list farms by farmer ID: %w", err)
	}
	defer rows.Close()

	var farms []*models.Farm
	for rows.Next() {
		farm := &models.Farm{}
		err = rows.Scan(
			&farm.ID,
			&farm.FarmerID,
			&farm.Name,
			&farm.Address,
			&farm.GeoLoc,
			&farm.Size,
			&farm.CropTypes,
			&farm.IsVerified,
			&farm.CreatedAt,
			&farm.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan farm: %w", err)
		}
		farms = append(farms, farm)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return farms, nil
}
