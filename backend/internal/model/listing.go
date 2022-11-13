package model

type Listing struct {
	ID     int    `form:"id" json:"id" db:"id"`
	Name   string `form:"name" json:"name" db:"name" binding:"required"`
	UserID int    `form:"user_id" json:"user_id" db:"user_id" binding:"required"`
}
