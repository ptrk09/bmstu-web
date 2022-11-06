package service

import "web/internal/repository"

type CalendarService struct {
	repo repository.Calendar
}

func NewCalendarService(repo repository.Calendar) *CalendarService {
	return &CalendarService{repo}
}
