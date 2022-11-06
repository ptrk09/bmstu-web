package model

type Listing struct {
	ID            int     `json:"id" db:"id"`
	Name          string  `json:"name" binding:"required"`
	UserID        int     `json:"listingID" db:"listing_id" binding:"required"`
	Neighbourhood string  `json:"neighbourhood" binding:"required"`
	ApartType     string  `json:"apartType" binding:"required"`
	Price         float64 `json:"price" binding:"required"`
	MinNights     int     `json:"minNights" db:"minimum_nights" binding:"required"`
}
