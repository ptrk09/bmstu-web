package model

type User struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UpdateUserInput struct {
	Name     string `json:"name" db:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
