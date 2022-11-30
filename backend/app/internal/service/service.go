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
	GetListings(id int, name string, user_id int) ([]model.Listing, error)
	CreateListing(listing model.Listing) (int, error)
	UpdateListing(id int, name string) (int, error)
	DeleteListing(id int) (int, error)
}

type ListingDetailed interface {
	GetListingsDetailed(
		id int,
		listing_id int,
		description string,
		neighbourhood string,
		apartTypeId int,
		price float32,
		minimumNights int,
	) ([]model.ListingDetailed, error)
	CreateListingDetailed(listingDetailed model.ListingDetailed) (int, error)
	UpdateListingDetailed(
		id int,
		description string,
		price float32,
		minimumNights int,
	) (int, error)
	DeleteListingDetailed(id int) (int, error)
}

type ListingImage interface {
	GetListingImages(id int, listingId int) ([]model.ListingImage, error)
	CreateListingImage(listingImage model.ListingImage) (int, error)
	UpdateLisingImage(id int, imagePath string) (int, error)
	DeleteListingImage(id int) (int, error)
}

type Calendar interface {
	GetAllCalendarInfo() ([]model.Calendar, error)
	CreateCalendarInfo(calendarInfo model.Calendar) (int, error)
	UpdateCalendarInfo(calendarInfo model.Calendar) (int, error)
	DeleteCalendarInfo(calendarId int) (int, error)
}

type Booking interface {
	GetBookings() ([]model.Booking, error)
	CreateBooking(booking model.Booking) (int, error)
	UpdateBooking(booking model.Booking) (int, error)
	DeleteBooking(bookingId int) (int, error)
}

type Service struct {
	Authorization
	User
	Listing
	ListingDetailed
	ListingImage
	Calendar
	Booking
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization:   NewAuthService(repo.Authorization),
		User:            NewUserService(repo.User),
		Listing:         NewListingService(repo.Listing),
		ListingDetailed: NewListingDetailedService(repo.ListingDetailed),
		ListingImage:    NewListingImageService(repo.ListingImage),
		Calendar:        NewCalendarService(repo.Calendar),
		Booking:         NewBookingService(repo.Booking),
	}
}
