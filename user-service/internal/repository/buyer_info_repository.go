package repository

import (
	"database/sql"
	"fmt"
	"user-service/internal/models"
)

type BuyerInfoStore struct {
	db *sql.DB
}

func NewBuyerInfoStore(db *sql.DB) *BuyerInfoStore {
	return &BuyerInfoStore{
		db: db,
	}
}

func (s *BuyerInfoStore) CreateBuyerInfo(buyerInfo *models.BuyerInfo) error {
	query := `INSERT INTO buyer_info (buyer_id, delivery_address, payment_method) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(query, buyerInfo.BuyerID, buyerInfo.DeliveryAddress, buyerInfo.PaymentMethod)

	if err != nil {
		return fmt.Errorf("could not create buyer info: %w", err)
	}

	fmt.Println("Buyer info created successfully")

	return nil
}

func (s *BuyerInfoStore) UpdateBuyerInfo(buyerInfo *models.BuyerInfo) error {
	query := `UPDATE buyer_info SET delivery_address = $1, payment_method = $2 WHERE buyer_id = $3`
	_, err := s.db.Exec(query, buyerInfo.DeliveryAddress, buyerInfo.PaymentMethod, buyerInfo.BuyerID)

	if err != nil {
		return fmt.Errorf("could not update buyer info: %w", err)
	}

	fmt.Println("Buyer info updated successfully")

	return nil
}

func (s *BuyerInfoStore) GetBuyerInfoByBuyerID(buyerID int) (*models.BuyerInfo, error) {
	query := `SELECT buyer_id, delivery_address, payment_method FROM buyer_info WHERE buyer_id = $1`
	row := s.db.QueryRow(query, buyerID)

	buyerInfo := &models.BuyerInfo{}
	err := row.Scan(
		&buyerInfo.BuyerID,
		&buyerInfo.DeliveryAddress,
		&buyerInfo.PaymentMethod,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not get buyer info: %w", err)
	}

	return buyerInfo, nil
}

func (s *BuyerInfoStore) DeleteBuyerInfo(buyerID int) error {
	query := `DELETE FROM buyer_info WHERE buyer_id = $1`
	_, err := s.db.Exec(query, buyerID)

	if err != nil {
		return fmt.Errorf("could not delete buyer info: %w", err)
	}

	fmt.Println("Buyer info deleted successfully")

	return nil
}
