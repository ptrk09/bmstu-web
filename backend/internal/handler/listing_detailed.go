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

// func (h *Handler) createListing(ctx *gin.Context) {
// 	var listing model.Listing

// 	if err := ctx.ShouldBindJSON(&listing); err != nil {
// 		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	listingId, err := h.services.Listing.CreateListing(listing)
// 	if err != nil {
// 		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"id": listingId,
// 	})
// }

// func (h *Handler) updateListing(ctx *gin.Context) {
// 	var listing model.Listing

// 	if err := ctx.ShouldBindJSON(&listing); err != nil {
// 		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	id := listing.ID
// 	name := listing.Name
// 	fmt.Print("id = ", id, "\n")
// 	fmt.Print("name = ", name, "\n")

// 	userId, err := h.services.Listing.UpdateListing(id, name)
// 	if err != nil {
// 		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"user_id": userId,
// 	})
// }

// func (h *Handler) deleteListing(ctx *gin.Context) {
// 	var listing model.Listing

// 	if err := ctx.ShouldBindJSON(&listing); err != nil {
// 		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	id := listing.ID

// 	listingId, err := h.services.Listing.DeleteListing(id)
// 	if err != nil {
// 		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"id": listingId,
// 	})
// }
