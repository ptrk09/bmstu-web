package model

import "time"

type Booking struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id" binding:"required"`
	ListingID int       `json:"listing_id" db:"listing_id" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
}
