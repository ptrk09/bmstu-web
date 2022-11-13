package service

import (
	"web/internal/model"
	"web/internal/repository"
)

type CalendarService struct {
	repo repository.Calendar
}

func NewCalendarService(repo repository.Calendar) *CalendarService {
	return &CalendarService{repo}
}

func (s *CalendarService) GetAllCalendarInfo() ([]model.Calendar, error) {
	return s.repo.GetAllCalendarInfo()
}

func (s *CalendarService) CreateCalendarInfo(calendarInfo model.Calendar) (int, error) {
	return s.repo.CreateCalendarInfo(calendarInfo)
}

func (s *CalendarService) UpdateCalendarInfo(calendarInfo model.Calendar) (int, error) {
	return s.repo.UpdateCalendarInfo(calendarInfo)
}

func (s *CalendarService) DeleteCalendarInfo(calendarId int) (int, error) {
	return s.repo.DeleteCalendarInfo(calendarId)
}
