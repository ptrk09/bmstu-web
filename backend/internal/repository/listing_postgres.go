package repository

import (
	"fmt"
	"web/internal/model"

	"github.com/jmoiron/sqlx"
)

type ListingPostgres struct {
	db *sqlx.DB
}

func NewListingPostgres(db *sqlx.DB) *ListingPostgres {
	return &ListingPostgres{db}
}

func (r *ListingPostgres) GetListingsByName(name string) ([]model.Listing, error) {
	var listings []model.Listing

	query := fmt.Sprintf("SELECT * FROM %s", teamsTable)
	err := r.db.Select(&teams, query)

	return teams, err
}
