package handler

import (
	"net/http"
	"web/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type getAllListingsDetailedResponse struct {
	Data []model.ListingDetailed `json:"data"`
}

func (h *Handler) getListingsDetailed(ctx *gin.Context) {
	var listingDetailed model.ListingDetailed

	if err := ctx.ShouldBindWith(&listingDetailed, binding.Query); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := listingDetailed.ID
	listing_id := listingDetailed.ListingID
	description := listingDetailed.Description
	neighbourhood := listingDetailed.Neighbourhood
	apartTypeId := listingDetailed.ApartTypeId
	price := listingDetailed.Price
	minimumNights := listingDetailed.MinimumNights

	listingsDetailed, err := h.services.ListingDetailed.GetListingsDetailed(
		id,
		listing_id,
		description,
		neighbourhood,
		apartTypeId,
		price,
		minimumNights,
	)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllListingsDetailedResponse{
		Data: listingsDetailed,
	})
}

func (h *Handler) createListingDetailed(ctx *gin.Context) {
	var listingDetailed model.ListingDetailed

	if err := ctx.ShouldBindJSON(&listingDetailed); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	listingDetailedId, err := h.services.ListingDetailed.CreateListingDetailed(listingDetailed)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": listingDetailedId,
	})
}

func (h *Handler) updateListingDetailed(ctx *gin.Context) {
	var listingDetailed model.ListingDetailed

	if err := ctx.ShouldBindJSON(&listingDetailed); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := listingDetailed.ID
	description := listingDetailed.Description
	price := listingDetailed.Price
	minimim_nights := listingDetailed.MinimumNights

	listingDetailedId, err := h.services.ListingDetailed.UpdateListingDetailed(
		id,
		description,
		price,
		minimim_nights,
	)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": listingDetailedId,
	})
}

func (h *Handler) deleteListingDetailed(ctx *gin.Context) {
	var listingDetailed model.ListingDetailed

	if err := ctx.ShouldBindJSON(&listingDetailed); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := listingDetailed.ID

	listingDetailedId, err := h.services.ListingDetailed.DeleteListingDetailed(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": listingDetailedId,
	})
}
