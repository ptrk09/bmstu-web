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

	fmt.Print("name = ", name, "\n")
	query := fmt.Sprintf("SELECT * FROM %s WHERE name = %s", listingsTable, name)
	err := r.db.Select(&listings, query)
	fmt.Print("listings = ", listings, "\n")

	return listings, err
}
