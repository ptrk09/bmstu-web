package service

import (
	"web/internal/model"
	"web/internal/repository"
)

type ListingService struct {
	repo repository.Listing
}

func NewListingService(repo repository.Listing) *ListingService {
	return &ListingService{repo}
}

func (s *ListingService) CreateListing(model.Listing) (int, error) {
	return 0, nil
}

func (s *ListingService) GetListingsByName(id int) (model.Listing, error) {
	return model.Listing{}, nil
}

func (s *ListingService) GetAllListings() ([]model.Listing, error) {
	return []model.Listing{}, nil
}
