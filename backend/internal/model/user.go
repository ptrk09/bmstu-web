package model

type User struct {
	ID       int    `json:"id" db:"id"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	PLevel   int    `json:"plevel" binding:"required"`
}
