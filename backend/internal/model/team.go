package model

type Team struct {
	ID      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	OwnerId int    `json:"owner_id" db:"owner_id"`
}
