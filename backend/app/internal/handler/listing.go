package handler

import (
	"net/http"
	"strconv"
	"web/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type getAllListingsResponse struct {
	Data []model.Listing `json:"data"`
}

// getListings godoc
// @Summary Get listings
// @Tags listing
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListingsResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param id query int false "ID of a listing."
// @Param name query string false "Name of a listing."
// @Param user_id query int false "ID of a user."
// @Router /listing/ [get]
func (h *Handler) getListings(ctx *gin.Context) {
	var listing model.Listing

	if err := ctx.ShouldBindWith(&listing, binding.Query); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := listing.ID
	name := listing.Name
	userId := listing.UserID

	listings, err := h.services.Listing.GetListings(id, name, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllListingsResponse{
		Data: listings,
	})
}

// createListing godoc
// @Summary Create listing
// @Tags listing
// @Accept  json
// @Produce  json
// @Success 200 {object} responseWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param name body string true "Name of a listing."
// @Param user_id body int true "ID of a user."
// @Router /listing/ [post]
func (h *Handler) createListing(ctx *gin.Context) {
	var listing model.Listing

	if err := ctx.ShouldBindJSON(&listing); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if listing.Name == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "Parameter listing name is not defined.")
		return
	}

	if listing.UserID == 0 {
		newErrorResponse(ctx, http.StatusBadRequest, "Parameter user id is not defined.")
		return
	}

	listingId, err := h.services.Listing.CreateListing(listing)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, responseWithId{listingId})
}

// updateListing godoc
// @Summary Update listing
// @Tags listing
// @Accept  json
// @Produce  json
// @Success 200 {object} responseWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param id path int true "ID of a listing."
// @Param name body string true "Name of a listing."
// @Router /listing/{id} [patch]
func (h *Handler) updateListing(ctx *gin.Context) {
	var listing model.Listing

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := ctx.ShouldBindJSON(&listing); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	name := listing.Name

	userId, err := h.services.Listing.UpdateListing(id, name)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, responseWithId{userId})
}

// deleteListing godoc
// @Summary Delete listing
// @Tags listing
// @Accept  json
// @Produce  json
// @Success 200 {object} responseWithId
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Param Authorization header string true "Authorization"
// @Param id query int true "ID of a listing."
// @Router /listing/{id} [delete]
func (h *Handler) deleteListing(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	listingId, err := h.services.Listing.DeleteListing(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, responseWithId{listingId})
}
