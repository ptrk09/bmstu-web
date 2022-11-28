package model

type ListingDetailed struct {
	ID            int     `form:"id" json:"id" db:"id"`
	ListingID     int     `form:"listing_id" json:"listing_id" db:"listing_id"`
	Description   string  `form:"description" json:"description" db:"description"`
	Neighbourhood string  `form:"neighbourhood" json:"neighbourhood" db:"neighbourhood"`
	ApartTypeId   int     `form:"apart_type_id" json:"apart_type_id" db:"apart_type_id"`
	Price         float32 `form:"price" json:"price" db:"price"`
	MinimumNights int     `form:"minimum_nights" json:"minimum_nights" db:"minimum_nights"`
}
