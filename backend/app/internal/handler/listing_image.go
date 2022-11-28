package handler

import (
	"net/http"
	"strconv"
	"web/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type getAllListingsImagesResponse struct {
	Data []model.ListingImage `json:"data"`
}

// getListingImages godoc
// @Summary Get listing images
// @Tags listingImage
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListingsImagesResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param id query int false "ID of a listing image."
// @Param listing_id query int false "Listing ID of a listing image."
// @Router /listing_image/ [get]
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

// createListingImage godoc
// @Summary Create listing images
// @Tags listingImage
// @Accept  json
// @Produce  json
// @Success 200 {object} responseWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param listing_id body int true "Listing ID of a listing image."
// @Param image_path body string true "Image path of a listing image."
// @Router /listing_image/ [post]
func (h *Handler) createListingImage(ctx *gin.Context) {
	var listingImage model.ListingImage

	if err := ctx.ShouldBindWith(&listingImage, binding.Query); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if listingImage.ListingID == 0 {
		newErrorResponse(ctx, http.StatusBadRequest, "Parameter listing id is not defined.")
		return
	}

	if listingImage.ImagePath == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "Parameter image path is not defined.")
		return
	}

	listingImageId, err := h.services.ListingImage.CreateListingImage(listingImage)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, responseWithId{listingImageId})
}

// updateListingImage godoc
// @Summary Update listing images
// @Tags listingImage
// @Accept json
// @Produce json
// @Success 200 {object} responseWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param id path int false "ID of a listing image."
// @Param image_path body string false "Image path of a listing image."
// @Router /listing_image/{id} [patch]
func (h *Handler) updateListingImage(ctx *gin.Context) {
	var listingImage model.ListingImage

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := ctx.ShouldBindWith(&listingImage, binding.Query); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	image_path := listingImage.ImagePath

	listingImageId, err := h.services.ListingImage.UpdateLisingImage(id, image_path)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, responseWithId{listingImageId})
}

// deleteListingImage godoc
// @Summary Delete listing images
// @Tags listingImage
// @Accept json
// @Produce json
// @Success 200 {object} responseWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param id path int true "ID of a listing image."
// @Router /listing_image/{id} [delete]
func (h *Handler) deleteListingImage(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	listingImageId, err := h.services.ListingImage.DeleteListingImage(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, responseWithId{listingImageId})
}
