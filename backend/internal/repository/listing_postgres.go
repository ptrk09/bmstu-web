package repository

import "github.com/jmoiron/sqlx"

type ListingPostgres struct {
	db *sqlx.DB
}

func NewListingPostgres(db *sqlx.DB) *ListingPostgres {
	return &ListingPostgres{db}
}
