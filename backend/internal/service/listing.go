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

func (s *ListingService) CreateListing(listing model.Listing) (int, error) {
	return s.repo.CreateListing(listing)
}

func (s *ListingService) GetListings(id int, name string, user_id int) ([]model.Listing, error) {
	return s.repo.GetListings(id, name, user_id)
}
