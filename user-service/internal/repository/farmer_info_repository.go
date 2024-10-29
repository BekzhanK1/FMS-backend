package repository

import (
	"database/sql"
	"fmt"
	"user-service/internal/models"
)

type FarmerInfoStore struct {
	db *sql.DB
}

func NewFarmerInfoStore(db *sql.DB) *FarmerInfoStore {
	return &FarmerInfoStore{
		db: db,
	}
}

func (s *FarmerInfoStore) CreateFarmerInfo(farmerInfo *models.FarmerInfo) error {
	query := `INSERT INTO farmer_info (farmer_id, rating, experience, bio) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(query, farmerInfo.FarmerID, farmerInfo.Rating, farmerInfo.Experience, farmerInfo.Bio)

	if err != nil {
		return fmt.Errorf("could not create farmer info: %w", err)
	}

	fmt.Println("Farmer info created successfully")

	return nil
}

func (s *FarmerInfoStore) UpdateFarmerInfo(farmerInfo *models.FarmerInfo) error {
	query := `UPDATE farmer_info SET rating = $1, experience = $2, bio = $3 WHERE farmer_id = $5`
	_, err := s.db.Exec(query, farmerInfo.Rating, farmerInfo.Experience, farmerInfo.Bio, farmerInfo.FarmerID)

	if err != nil {
		return fmt.Errorf("could not update farmer info: %w", err)
	}

	fmt.Println("Farmer info updated successfully")

	return nil
}

func (s *FarmerInfoStore) GetFarmerInfoByFarmerId(farmerId int) (*models.FarmerInfo, error) {
	query := `SELECT farmer_id, rating, experience, bio FROM farmer_info WHERE farmer_id = $1`
	row := s.db.QueryRow(query, farmerId)

	farmerInfo := &models.FarmerInfo{}
	err := row.Scan(
		&farmerInfo.FarmerID,
		&farmerInfo.Rating,
		&farmerInfo.Experience,
		&farmerInfo.Bio,
	)

	if err != nil {
		return nil, fmt.Errorf("could not get farmer info: %w", err)
	}

	return farmerInfo, nil
}

func (s *FarmerInfoStore) DeleteFarmerInfo(farmerId int) error {
	query := `DELETE FROM farmer_info WHERE farmer_id = $1`
	_, err := s.db.Exec(query, farmerId)

	if err != nil {
		return fmt.Errorf("could not delete farmer info: %w", err)
	}

	fmt.Println("Farmer info deleted successfully")

	return nil
}