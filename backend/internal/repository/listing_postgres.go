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

func (r *ListingPostgres) GetListings(id int, name string, user_id int) ([]model.Listing, error) {
	var listings []model.Listing

	// query := fmt.Sprintf("SELECT * FROM %s WHERE id = %d AND name = %s AND user_id = %d;", listingsTable, id, name, user_id)
	query := fmt.Sprintf("SELECT * FROM %s WHERE", listingsTable)

	if id == 0 {
		query = query + fmt.Sprintf(" id = %s", "id")
	} else {
		query = query + fmt.Sprintf(" id = %d", id)
	}

	if name == "" {
		query = query + fmt.Sprintf(" name = %s", "name")
	} else {
		query = query + fmt.Sprintf(" AND name = %s", name)
	}

	if user_id == 0 {
		query = query + fmt.Sprintf(" user_id = %s", "user_id")
	} else {
		query = query + fmt.Sprintf(" AND user_id = %d", user_id)
	}

	fmt.Print(query)

	err := r.db.Select(&listings, query)

	return listings, err
}
