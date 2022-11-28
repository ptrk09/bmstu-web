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

type Repository struct {
	Authorization
	User
	Listing
	ListingDetailed
	ListingImage
	Calendar
	Booking
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:   NewAuthPostgres(db),
		User:            NewUserPostgres(db),
		Listing:         NewListingPostgres(db),
		ListingDetailed: NewListingDetailedPostgres(db),
		ListingImage:    NewListingImagePostgres(db),
		Calendar:        NewCalendarPostgres(db),
		Booking:         NewBookingPostgres(db),
	}
}
