package service

import (
	"web/internal/model"
	"web/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
}

type Player interface {
	// Create(userId int, player model.Player) (int, error)
}

type Team interface {
	Create(userId int, team model.Team) (int, error)
	GetAll(userId int) ([]model.Team, error)
}

type Service struct {
	Authorization
	User
	Player
	Team
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Team:          NewTeamService(repo.Team),
	}
}
