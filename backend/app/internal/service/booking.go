package service

import (
	"web/internal/model"
	"web/internal/repository"
)

type BookingService struct {
	repo repository.Booking
}

func NewBookingService(repo repository.Booking) *BookingService {
	return &BookingService{repo}
}

func (s *BookingService) GetBookings() ([]model.Booking, error) {
	return s.repo.GetBookings()
}

func (s *BookingService) CreateBooking(booking model.Booking) (int, error) {
	return s.repo.CreateBooking(booking)
}

func (s *BookingService) UpdateBooking(booking model.Booking) (int, error) {
	return s.repo.UpdateBooking(booking)
}

func (s *BookingService) DeleteBooking(bookingId int) (int, error) {
	return s.repo.DeleteBooking(bookingId)
}
