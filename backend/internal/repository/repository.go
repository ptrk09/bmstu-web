package repository

import (
	"web/internal/model"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(login, password string) (model.User, error)
}

type User interface {
}

type Player interface {
}

type Team interface {
	Create(userId int, team model.Team) (int, error)
	GetAll(userId int) ([]model.Team, error)
}

type Repository struct {
	Authorization
	User
	Player
	Team
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Team:          NewTeamPostgres(db),
	}
}
