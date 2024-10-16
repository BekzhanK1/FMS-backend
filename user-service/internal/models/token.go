package models

import "time"

type Token struct {
	ID         int       `db:"id"`
	UserID     int    	 `db:"userId"`
	Token      string    `db:"token"`
	Expiration time.Time `db:"expiration"`
	UpdatedAt  time.Time `db:"updated_at"`
}
