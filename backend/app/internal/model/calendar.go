package model

import "time"

type Calendar struct {
	ID        int       `json:"id" db:"id"`
	ListingID int       `json:"listing_id" db:"listing_id" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	Available bool      `json:"available" binding:"required"`
}
