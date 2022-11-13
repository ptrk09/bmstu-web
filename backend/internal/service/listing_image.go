package service

import (
	"web/internal/model"
	"web/internal/repository"
)

type ListingImageService struct {
	repo repository.ListingImage
}

func NewListingImageService(repo repository.ListingImage) *ListingImageService {
	return &ListingImageService{repo}
}

func (s *ListingImageService) GetListingImages(id int, listingId int) ([]model.ListingImage, error) {
	return s.repo.GetListingImages(id, listingId)
}

func (s *ListingImageService) CreateListingImage(listingImage model.ListingImage) (int, error) {
	return s.repo.CreateListingImage(listingImage)
}

func (s *ListingImageService) UpdateLisingImage(id int, imagePath string) (int, error) {
	return s.repo.UpdateLisingImage(id, imagePath)
}

func (s *ListingImageService) DeleteListingImage(id int) (int, error) {
	return s.repo.DeleteListingImage(id)
}
