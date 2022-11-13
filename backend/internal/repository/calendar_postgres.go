package repository

import (
	"fmt"
	"web/internal/model"

	"github.com/jmoiron/sqlx"
)

type CalendarPostgres struct {
	db *sqlx.DB
}

func NewCalendarPostgres(db *sqlx.DB) *CalendarPostgres {
	return &CalendarPostgres{db}
}

func (r *CalendarPostgres) GetAllCalendarInfo() ([]model.Calendar, error) {
	var calendar []model.Calendar

	query := fmt.Sprintf(
		`SELECT * FROM %s`,
		calendarTable,
	)
	err := r.db.Select(&calendar, query)

	return calendar, err
}

func (r *CalendarPostgres) CreateCalendarInfo(calendarInfo model.Calendar) (int, error) {
	var calendarInfoId int
	query := fmt.Sprintf(
		`INSERT INTO %s (listing_id, date, available) 
		VALUES ($1, $2, $3) RETURNING id`,
		listingsTable,
	)

	row := r.db.QueryRow(
		query,
		calendarInfo.ListingID,
		calendarInfo.Date,
		// calendarInfo.Date.Time,
		calendarInfo.Available,
	)
	if err := row.Scan(&calendarInfoId); err != nil {
		return 0, err
	}

	return calendarInfoId, nil
}

func (r *CalendarPostgres) UpdateCalendarInfo(calendarInfo model.Calendar) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var calendarInfoId int
	updateUserQuery := fmt.Sprintf(
		`UPDATE %s
		SET listing_id = $1, date = $2, available = $3
		WHERE id = $4
		RETURNING id;`,
		listingsTable,
	)

	row := tx.QueryRow(
		updateUserQuery,
		calendarInfo.ListingID,
		calendarInfo.Date,
		// calendarInfo.Date.Time,
		calendarInfo.Available,
		calendarInfo.ID,
	)
	if err := row.Scan(&calendarInfoId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return calendarInfoId, tx.Commit()
}

func (r *CalendarPostgres) DeleteCalendarInfo(calendarId int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var deleteRowId int
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", calendarTable)

	row := tx.QueryRow(query, calendarId)
	err = row.Scan(&deleteRowId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return deleteRowId, tx.Commit()
}
