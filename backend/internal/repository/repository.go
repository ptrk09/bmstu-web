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
	GetAllUsers() ([]model.User, error)
	GetUserById(userId int) (model.User, error)
	UpdateUser(userId int, input model.UpdateUserInput) (int, error)
	DeleteUser(userId int) (int, error)
}

type Listing interface {
	GetListings(id int, name string, user_id int) ([]model.Listing, error)
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
		User:          NewUserPostgres(db),
		// Team:          NewTeamPostgres(db),
	}
}
