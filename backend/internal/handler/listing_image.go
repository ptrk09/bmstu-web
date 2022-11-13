package handler

import (
	"net/http"
	"web/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type getAllListingsImagesResponse struct {
	Data []model.ListingImage `json:"data"`
}

func (h *Handler) getListingImages(ctx *gin.Context) {
	var listingImage model.ListingImage

	if err := ctx.ShouldBindWith(&listingImage, binding.Query); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := listingImage.ID
	listing_id := listingImage.ListingID

	listingImages, err := h.services.ListingImage.GetListingImages(id, listing_id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllListingsImagesResponse{
		Data: listingImages,
	})
}

func (h *Handler) createListingImage(ctx *gin.Context) {
	var listingImage model.ListingImage

	if err := ctx.ShouldBindWith(&listingImage, binding.Query); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	listingImageId, err := h.services.ListingImage.CreateListingImage(listingImage)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": listingImageId,
	})
}

func (h *Handler) updateListingImage(ctx *gin.Context) {
	var listingImage model.ListingImage

	if err := ctx.ShouldBindWith(&listingImage, binding.Query); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := listingImage.ID
	image_path := listingImage.ImagePath

	listingImageId, err := h.services.ListingImage.UpdateLisingImage(id, image_path)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": listingImageId,
	})
}

func (h *Handler) deleteListingImage(ctx *gin.Context) {
	var listingImage model.ListingImage

	if err := ctx.ShouldBindWith(&listingImage, binding.Query); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := listingImage.ID

	listingImageId, err := h.services.ListingImage.DeleteListingImage(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": listingImageId,
	})
}
