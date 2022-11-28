package model

type ListingImage struct {
	ID        int    `form:"id" json:"id" db:"id"`
	ListingID int    `form:"listing_id" json:"listing_id" db:"listing_id"`
	ImagePath string `form:"image_path" json:"image_path" db:"image_path"`
}
