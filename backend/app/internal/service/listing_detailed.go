package service

import (
	"web/internal/model"
	"web/internal/repository"
)

type ListingDetailedService struct {
	repo repository.ListingDetailed
}

func NewListingDetailedService(repo repository.ListingDetailed) *ListingDetailedService {
	return &ListingDetailedService{repo}
}

func (s *ListingDetailedService) GetListingsDetailed(
	id int,
	listing_id int,
	description string,
	neighbourhood string,
	apartTypeId int,
	price float32,
	minimumNights int,
) ([]model.ListingDetailed, error) {
	return s.repo.GetListingsDetailed(
		id,
		listing_id,
		description,
		neighbourhood,
		apartTypeId,
		price,
		minimumNights,
	)
}

func (s *ListingDetailedService) CreateListingDetailed(listingDetailed model.ListingDetailed) (int, error) {
	return s.repo.CreateListingDetailed(listingDetailed)
}

func (s *ListingDetailedService) UpdateListingDetailed(
	id int,
	description string,
	price float32,
	minimumNights int,
) (int, error) {
	return s.repo.UpdateListingDetailed(
		id,
		description,
		price,
		minimumNights,
	)
}

func (s *ListingDetailedService) DeleteListingDetailed(id int) (int, error) {
	return s.repo.DeleteListingDetailed(id)
}
