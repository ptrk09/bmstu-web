package service

import (
	"web/internal/model"
	"web/internal/repository"
)

type TeamService struct {
	repo repository.Team
}

func NewTeamService(repo repository.Team) *TeamService {
	return &TeamService{repo}
}

func (s *TeamService) Create(userId int, team model.Team) (int, error) {
	return s.repo.Create(userId, team)
}

func (s *TeamService) GetAll(userId int) ([]model.Team, error) {
	return s.repo.GetAll(userId)
}
