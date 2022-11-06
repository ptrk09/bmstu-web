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
}

type Listing interface {
	CreateListing(model.Listing) (int, error)
	GetListingById(id int) (model.Listing, error)
	GetAllListings() ([]model.Listing, error)
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
		Listing:       NewListingService(repo.Listing),
		Calendar:      NewCalendarService(repo.Calendar),
		// Team:          NewTeamService(repo.Team),
	}
}
