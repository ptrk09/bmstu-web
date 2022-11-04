package service

import (
	"web/internal/repository"
)

type PlayerService struct {
	repo repository.Player
}

func NewPlayerService(repo repository.Player) *PlayerService {
	return &PlayerService{repo}
}

// func (s *PlayerService) Create(userId int, player model.Player) (int, error) {

// }
