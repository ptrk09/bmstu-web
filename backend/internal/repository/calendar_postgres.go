package repository

import "github.com/jmoiron/sqlx"

type CalendarPostgres struct {
	db *sqlx.DB
}

func NewCalendarPostgres(db *sqlx.DB) *CalendarPostgres {
	return &CalendarPostgres{db}
}
