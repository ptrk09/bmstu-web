package repository

import (
	"fmt"
	"time"
	"web/internal/model"

	"github.com/jmoiron/sqlx"
)

type BookingPostgres struct {
	db *sqlx.DB
}

func NewBookingPostgres(db *sqlx.DB) *BookingPostgres {
	return &BookingPostgres{db}
}

func (r *BookingPostgres) GetBookings() ([]model.Booking, error) {
	bookings := []model.Booking{}
	// var rawData []byte

	query := fmt.Sprintf(
		`SELECT * FROM %s`,
		bookingsTable,
	)
	err := r.db.Select(&bookings, query)
	// fmt.Println(bookings)

	return bookings, err
}

func (r *BookingPostgres) CreateBooking(booking model.Booking) (int, error) {
	var bookingId int
	query := fmt.Sprintf(
		`INSERT INTO %s (user_id, listing_id, date)
		VALUES ($1, $2, $3) RETURNING id`,
		bookingsTable,
	)

	row := r.db.QueryRow(
		query,
		booking.UserID,
		booking.ListingID,
		// booking.Date.Time,
		time.Time(booking.Date),
	)
	if err := row.Scan(&bookingId); err != nil {
		return 0, err
	}

	return bookingId, nil
}

func (r *BookingPostgres) UpdateBooking(booking model.Booking) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var calendarInfoId int
	updateUserQuery := fmt.Sprintf(
		`UPDATE %s
		SET user_id = $1, listing_id = $2, date = $3
		WHERE id = $4
		RETURNING id;`,
		bookingsTable,
	)

	row := tx.QueryRow(
		updateUserQuery,
		booking.UserID,
		booking.ListingID,
		// booking.Date,
		booking.Date,
		booking.ID,
	)
	if err := row.Scan(&calendarInfoId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return calendarInfoId, tx.Commit()
}

func (r *BookingPostgres) DeleteBooking(bookingId int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var deleteRowId int
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", bookingsTable)

	row := tx.QueryRow(query, bookingId)
	err = row.Scan(&deleteRowId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return deleteRowId, tx.Commit()
}
