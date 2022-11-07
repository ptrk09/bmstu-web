package service

import (
	"web/internal/model"
	"web/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (UserContextData, error)
}

type User interface {
	GetAllUsers() ([]model.User, error)
	GetUserById(userId int) (model.User, error)
	DeleteUser(userId int) (int, error)
	UpdateUser(userId int, input model.UpdateUserInput) (int, error)
}

type Listing interface {
	CreateListing(model.Listing) (int, error)
	GetAllListings() ([]model.Listing, error)
	GetListings(id int, name string, user_id int) ([]model.Listing, error)
}

type Calendar interface{}

// type Team interface {
// 	Create(userId int, team model.Team) (int, error)
// 	GetAll(userId int) ([]model.Team, error)
// }

type Service struct {
	Authorization
	User
	Listing
	Calendar
	// Team
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		User:          NewUserService(repo.User),
		Listing:       NewListingService(repo.Listing),
		Calendar:      NewCalendarService(repo.Calendar),
		// Team:          NewTeamService(repo.Team),
	}
}
