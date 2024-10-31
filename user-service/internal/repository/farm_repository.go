package repository

import (
	"database/sql"
	"fmt"
	"user-service/internal/models"
	"user-service/types"
)

type FarmStore struct {
	db *sql.DB
}

func NewFarmStore(db *sql.DB) *FarmStore {
	return &FarmStore{
		db: db,
	}
}

func (s *FarmStore) CreateFarm(farm *models.Farm) (*models.Farm, error) {
	countQuery := `SELECT COUNT(*) FROM farms WHERE farmer_id = $1`
	var count int
	err := s.db.QueryRow(countQuery, farm.FarmerID).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("could not count farms: %w", err)
	}

	if count >= 5 {
		return nil, fmt.Errorf("farmer with ID %d has reached the maximum number of farms, which is 5", farm.FarmerID)
	}

	query := `INSERT INTO farms (name, address, farmer_id, geo_loc, size, crop_types, is_verified, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err = s.db.QueryRow(query, farm.Name, farm.Address, farm.FarmerID, farm.GeoLoc, farm.Size, farm.CropTypes, farm.IsVerified, farm.CreatedAt, farm.UpdatedAt).Scan(&farm.ID)
	if err != nil {
		return nil, fmt.Errorf("could not create farm: %w", err)
	}

	fmt.Println("Farm created successfully")
	return farm, nil
}

func (s *FarmStore) GetFarmByID(id int) (*types.FarmResponse, error) {
	query := `
	SELECT  
    f.id AS farm_id,
    f.name AS farm_name,
    f.address AS farm_address,
    f.geo_loc AS farm_geo_loc,
    f.size AS farm_size,
    f.crop_types AS farm_crop_types,
	
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
    fi.bio
	FROM farms f 
	JOIN users u ON f.farmer_id = u.id
	LEFT JOIN farmer_info fi ON fi.farmer_id = u.id
	WHERE f.id = $1 AND f.is_verified = true;
	`
	row := s.db.QueryRow(query, id)

	farm := &types.FarmResponse{}
	err := row.Scan(
		&farm.ID,
		&farm.Name,
		&farm.Address,
		&farm.GeoLoc,
		&farm.Size,
		&farm.CropTypes,
		&farm.Farmer.ID,
		&farm.Farmer.FirstName,
		&farm.Farmer.LastName,
		&farm.Farmer.Username,
		&farm.Farmer.Email,
		&farm.Farmer.Phone,
		&farm.Farmer.ProfilePicture,
		&farm.Farmer.Role,
		&farm.Farmer.Rating,
		&farm.Farmer.Experience,
		&farm.Farmer.Bio,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("could not get farm by ID: %w", err)
	}
	return farm, nil
}

func (s *FarmStore) ListFarms() ([]*types.FarmResponse, error) {
	query := `	
	SELECT  
    f.id AS farm_id,
    f.name AS farm_name,
    f.address AS farm_address,
    f.geo_loc AS farm_geo_loc,
    f.size AS farm_size,
    f.crop_types AS farm_crop_types,
	
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
    fi.bio
	FROM farms f 
	JOIN users u ON f.farmer_id = u.id
	LEFT JOIN farmer_info fi ON fi.farmer_id = u.id
	WHERE f.is_verified = true;
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not list farms: %w", err)
	}
	defer rows.Close()

	farms := make([]*types.FarmResponse, 0)
	for rows.Next() {
		farm := &types.FarmResponse{}
		err = rows.Scan(
			&farm.ID,
			&farm.Name,
			&farm.Address,
			&farm.GeoLoc,
			&farm.Size,
			&farm.CropTypes,
			&farm.Farmer.ID,
			&farm.Farmer.FirstName,
			&farm.Farmer.LastName,
			&farm.Farmer.Username,
			&farm.Farmer.Email,
			&farm.Farmer.Phone,
			&farm.Farmer.ProfilePicture,
			&farm.Farmer.Role,
			&farm.Farmer.Rating,
			&farm.Farmer.Experience,
			&farm.Farmer.Bio,
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

func (s *FarmStore) ListFarmsByFarmerID(farmerID int) ([]*types.FarmResponse, error) {
	query := `
	SELECT  
    f.id AS farm_id,
    f.name AS farm_name,
    f.address AS farm_address,
    f.geo_loc AS farm_geo_loc,
    f.size AS farm_size,
    f.crop_types AS farm_crop_types,
	
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
    fi.bio
	FROM farms f 
	JOIN users u ON f.farmer_id = u.id
	LEFT JOIN farmer_info fi ON fi.farmer_id = u.id
	WHERE f.farmer_id = $1 AND f.is_verified = true
	`
	rows, err := s.db.Query(query, farmerID)
	if err != nil {
		return nil, fmt.Errorf("could not list farms: %w", err)
	}
	defer rows.Close()

	farms := make([]*types.FarmResponse, 0)
	for rows.Next() {
		farm := &types.FarmResponse{}
		err = rows.Scan(
			&farm.ID,
			&farm.Name,
			&farm.Address,
			&farm.GeoLoc,
			&farm.Size,
			&farm.CropTypes,
			&farm.Farmer.ID,
			&farm.Farmer.FirstName,
			&farm.Farmer.LastName,
			&farm.Farmer.Username,
			&farm.Farmer.Email,
			&farm.Farmer.Phone,
			&farm.Farmer.ProfilePicture,
			&farm.Farmer.Role,
			&farm.Farmer.Rating,
			&farm.Farmer.Experience,
			&farm.Farmer.Bio,
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
