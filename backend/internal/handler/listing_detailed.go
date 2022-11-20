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

// getListingsDetailed godoc
// @Summary Get listings detailed
// @Tags listingDetailed
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListingsDetailedResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param id query int false "ID of a listing detailed."
// @Param listing_id query int false "Listing ID of a listing detailed."
// @Param description query string false "Description of a listing."
// @Param neighbourhood query string false "Neighbourhood of a listing."
// @Param apart_type_id query int false "Apart type ID of a listing."
// @Param price query float32 false "Price of a listing."
// @Param minimum_nights query int false "Minimum nights of a listing."
// @Router /listing_detailed/ [get]
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

// createListingsDetailed godoc
// @Summary Create listings detailed
// @Tags listingDetailed
// @Accept  json
// @Produce  json
// @Success 200 {object} responseWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param listing_id body int true "Listing ID of a listing detailed."
// @Param description body string false "Description of a listing."
// @Param neighbourhood body string false "Neighbourhood of a listing."
// @Param apart_type_id body string false "Apart type ID of a listing."
// @Param price body string true "Price of a listing."
// @Param minimum_nights body float32 false "Minimum nights of a listing."
// @Router /listing_detailed/ [post]
func (h *Handler) createListingDetailed(ctx *gin.Context) {
	var listingDetailed model.ListingDetailed

	if err := ctx.ShouldBindJSON(&listingDetailed); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if listingDetailed.ListingID == 0 {
		newErrorResponse(ctx, http.StatusBadRequest, "Parameter listing id is not defined.")
		return
	}

	if listingDetailed.Price == 0.0 {
		newErrorResponse(ctx, http.StatusBadRequest, "Parameter price is not defined.")
		return
	}

	listingDetailedId, err := h.services.ListingDetailed.CreateListingDetailed(listingDetailed)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, responseWithId{listingDetailedId})
}

// updateListingDetailed godoc
// @Summary Update listings detailed
// @Tags listingDetailed
// @Accept  json
// @Produce  json
// @Success 200 {object} responseWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param id query int false "ID of a listing detailed."
// @Param description query string false "Description of a listing."
// @Param price query float32 true "Price of a listing."
// @Param minimum_nights query int false "Minimum nights of a listing."
// @Router /listing_detailed/ [patch]
func (h *Handler) updateListingDetailed(ctx *gin.Context) {
	var listingDetailed model.ListingDetailed

	if err := ctx.ShouldBindWith(&listingDetailed, binding.Query); err != nil {
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

	ctx.JSON(http.StatusOK, responseWithId{listingDetailedId})
}

// deleteListingDetailed godoc
// @Summary Delete listings detailed
// @Tags listingDetailed
// @Accept  json
// @Produce  json
// @Success 200 {object} responseWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param id query int true "ID of a listing detailed."
// @Router /listing_detailed/ [delete]
func (h *Handler) deleteListingDetailed(ctx *gin.Context) {
	var listingDetailed model.ListingDetailed

	if err := ctx.ShouldBindWith(&listingDetailed, binding.Query); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := listingDetailed.ID

	listingDetailedId, err := h.services.ListingDetailed.DeleteListingDetailed(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, responseWithId{listingDetailedId})
}
