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

func (r *ListingPostgres) GetListings(id int, name string, userId int) ([]model.Listing, error) {
	var listings []model.Listing

	query := fmt.Sprintf("SELECT * FROM %s WHERE", listingsTable)

	if id == 0 {
		query = query + fmt.Sprintf(" id = %s", "id")
	} else {
		query = query + fmt.Sprintf(" id = %d", id)
	}

	if name == "" {
		query = query + fmt.Sprintf(" AND name = %s", "name")
	} else {
		query = query + fmt.Sprintf(" AND name = %s", name)
	}

	if userId == 0 {
		query = query + fmt.Sprintf(" AND user_id = %s", "user_id")
	} else {
		query = query + fmt.Sprintf(" AND user_id = %d", userId)
	}

	fmt.Print(query)

	err := r.db.Select(&listings, query)

	return listings, err
}

func (r *ListingPostgres) CreateListing(listing model.Listing) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listingId int
	query := fmt.Sprintf("INSERT INTO %s (name, user_id) values ('%s', %d) RETURNING id", listingsTable, listing.Name, listing.UserID)
	fmt.Println(query)

	row := tx.QueryRow(query)
	err = row.Scan(&listingId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listingId, tx.Commit()
}
