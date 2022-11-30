package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	typeDB = "postgres"

	usersTable           = "users"
	teamsTable           = "teams"
	rolesTable           = "roles"
	listingsTable        = "listings"
	lisingsDetailedTable = "listings_detailed"
	calendarTable        = "calendar"
	bookingsTable        = "bookings"
	listingsImagesTable  = "listings_images"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostrgersDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(typeDB, makeConfigString(cfg))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func makeConfigString(cfg Config) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode,
	)
}
