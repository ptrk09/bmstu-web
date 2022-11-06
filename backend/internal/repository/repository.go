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

type Listing interface {
}

type Calendar interface {
}

// type Team interface {
// 	Create(userId int, team model.Team) (int, error)
// 	GetAll(userId int) ([]model.Team, error)
// }

type Repository struct {
	Authorization
	User
	Listing
	Calendar
	// Team
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Listing:       NewListingPostgres(db),
		Calendar:      NewCalendarPostgres(db),
		// Team:          NewTeamPostgres(db),
	}
}
