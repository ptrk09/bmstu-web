package repository

import (
	"fmt"
	"web/internal/model"

	"github.com/jmoiron/sqlx"
)

type ListingDetailedPostgres struct {
	db *sqlx.DB
}

func NewListingDetailedPostgres(db *sqlx.DB) *ListingDetailedPostgres {
	return &ListingDetailedPostgres{db}
}

func (r *ListingDetailedPostgres) GetListingsDetailed(
	id int,
	listing_id int,
	description string,
	neighbourhood string,
	apartTypeId int,
	price float32,
	minimumNights int,
) ([]model.ListingDetailed, error) {
	var listingsDetailed []model.ListingDetailed

	query := fmt.Sprintf("SELECT * FROM %s WHERE", lisingsDetailedTable)

	if id == 0 {
		query = query + fmt.Sprintf(" id = %s", "id")
	} else {
		query = query + fmt.Sprintf(" id = %d", id)
	}

	if listing_id == 0 {
		query = query + fmt.Sprintf(" AND listing_id = %s", "listing_id")
	} else {
		query = query + fmt.Sprintf(" AND listing_id = %d", listing_id)
	}

	if description == "" {
		query = query + fmt.Sprintf(" AND description = %s", "description")
	} else {
		query = query + fmt.Sprintf(" AND description = %s", description)
	}

	if neighbourhood == "" {
		query = query + fmt.Sprintf(" AND neighbourhood = %s", "neighbourhood")
	} else {
		query = query + fmt.Sprintf(" AND neighbourhood = %s", neighbourhood)
	}

	if apartTypeId == 0 {
		query = query + fmt.Sprintf(" AND apart_type_id = %s", "apart_type_id")
	} else {
		query = query + fmt.Sprintf(" AND apart_type_id = %d", apartTypeId)
	}

	if price == 0.0 {
		query = query + fmt.Sprintf(" AND price = %s", "price")
	} else {
		query = query + fmt.Sprintf(" AND price = %f", price)
	}

	if minimumNights == 0 {
		query = query + fmt.Sprintf(" AND minimum_nights = %s", "minimum_nights")
	} else {
		query = query + fmt.Sprintf(" AND minimum_nights = %d", minimumNights)
	}

	fmt.Print(query)

	err := r.db.Select(&listingsDetailed, query)

	return listingsDetailed, err
}

func (r *ListingDetailedPostgres) CreateListingDetailed(
	listingDetailed model.ListingDetailed,
) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listingId int
	query := fmt.Sprintf("INSERT INTO %s (listing_id, description, neighbourhood, apart_type_id, price, minimum_nights) values (%d, '%s', '%s', %d, %f, %d) RETURNING listing_id",
		listingsTable,
		listingDetailed.ListingID,
		listingDetailed.Description,
		listingDetailed.Neighbourhood,
		listingDetailed.ApartTypeId,
		listingDetailed.Price,
		listingDetailed.MinimumNights,
	)

	row := tx.QueryRow(query)
	err = row.Scan(&listingId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listingId, tx.Commit()
}

func (r *ListingDetailedPostgres) UpdateListingDetailed(
	id int,
	listing_id int,
	description string,
	neighbourhood string,
	apartTypeId int,
	price float32,
	minimumNights int,
) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listingId int
	query := fmt.Sprintf("UPDATE %s SET", lisingsDetailedTable)

	if listing_id == 0 {
		query = query + fmt.Sprintf(" listing_id = %s", "listing_id")
	} else {
		query = query + fmt.Sprintf(" listing_id = %d", listing_id)
	}

	if description == "" {
		query = query + fmt.Sprintf(", description = %s", "description")
	} else {
		query = query + fmt.Sprintf(", description = %s", description)
	}

	if neighbourhood == "" {
		query = query + fmt.Sprintf(", neighbourhood = %s", "neighbourhood")
	} else {
		query = query + fmt.Sprintf(", neighbourhood = %s", neighbourhood)
	}

	if apartTypeId == 0 {
		query = query + fmt.Sprintf(", apart_type_id = %s", "apart_type_id")
	} else {
		query = query + fmt.Sprintf(", apart_type_id = %d", apartTypeId)
	}

	if price == 0.0 {
		query = query + fmt.Sprintf(", price = %s", "price")
	} else {
		query = query + fmt.Sprintf(", price = %f", price)
	}

	if minimumNights == 0 {
		query = query + fmt.Sprintf(", minimum_nights = %s", "minimum_nights")
	} else {
		query = query + fmt.Sprintf(", minimum_nights = %d", minimumNights)
	}

	query = query + fmt.Sprintf(" WHERE id = %d RETURNING id", id)

	fmt.Print(query)

	row := tx.QueryRow(query)
	err = row.Scan(&listingId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listingId, tx.Commit()
}

func (r *ListingDetailedPostgres) DeleteListingDetailed(id int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listingDetailedId int
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d RETURNING id", lisingsDetailedTable, id)

	row := tx.QueryRow(query)
	err = row.Scan(&listingDetailedId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listingDetailedId, tx.Commit()
}
