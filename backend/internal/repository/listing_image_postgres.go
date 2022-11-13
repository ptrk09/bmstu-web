package repository

import (
	"fmt"
	"web/internal/model"

	"github.com/jmoiron/sqlx"
)

type ListingImagePostgres struct {
	db *sqlx.DB
}

func NewListingImagePostgres(db *sqlx.DB) *ListingImagePostgres {
	return &ListingImagePostgres{db}
}

func (r *ListingImagePostgres) GetListingImages(id int, listingId int) ([]model.ListingImage, error) {
	var listingImages []model.ListingImage

	query := fmt.Sprintf("SELECT * FROM %s WHERE", listingsImagesTable)

	if id == 0 {
		query = query + fmt.Sprintf(" id = %s", "id")
	} else {
		query = query + fmt.Sprintf(" id = %d", id)
	}

	if listingId == 0 {
		query = query + fmt.Sprintf(" AND listing_id = %s", "listingId")
	} else {
		query = query + fmt.Sprintf(" AND listing_id = %d", listingId)
	}

	err := r.db.Select(&listingImages, query)

	return listingImages, err
}

func (r *ListingImagePostgres) CreateListingImage(listingImage model.ListingImage) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listingImageId int
	query := fmt.Sprintf("INSERT INTO %s (listing_id, image_path) values (%d, '%s') RETURNING id",
		listingsImagesTable,
		listingImage.ListingID,
		listingImage.ImagePath,
	)

	row := tx.QueryRow(query)
	err = row.Scan(&listingImageId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listingImageId, tx.Commit()
}

func (r *ListingImagePostgres) UpdateLisingImage(id int, imagePath string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listingImageId int
	query := fmt.Sprintf("UPDATE %s SET", listingsImagesTable)

	if imagePath == "" {
		query = query + fmt.Sprintf(" image_path = %s", "image_path")
	} else {
		query = query + fmt.Sprintf(" image_path = %s", imagePath)
	}

	query = query + fmt.Sprintf(" WHERE id = %d RETURNING id", id)

	row := tx.QueryRow(query)
	err = row.Scan(&listingImageId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listingImageId, tx.Commit()
}

func (r *ListingImagePostgres) DeleteListingImage(id int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listingImageId int
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d RETURNING id", listingsImagesTable, id)

	row := tx.QueryRow(query)
	err = row.Scan(&listingImageId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listingImageId, tx.Commit()
}
