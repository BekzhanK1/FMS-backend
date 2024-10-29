package models

type BuyerInfo struct {
	BuyerID         int    `json:"buyer_id" db:"buyer_id"`
	DeliveryAddress string `json:"delivery_address" db:"delivery_address"`
	PaymentMethod   string `json:"payment_method" db:"payment_method"`
}
