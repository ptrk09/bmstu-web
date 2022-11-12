package handler

import (
	"fmt"
	"net/http"
	"web/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type getAllListingsResponse struct {
	Data []model.Listing `json:"data"`
}

func (h *Handler) getListings(ctx *gin.Context) {
	var listing model.Listing

	if err := ctx.ShouldBindWith(&listing, binding.Query); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := listing.ID
	name := listing.Name
	userId := listing.UserID

	fmt.Print("id = ", id, "\n")
	fmt.Print("name = ", name, "\n")
	fmt.Print("userId = ", userId, "\n")

	listings, err := h.services.Listing.GetListings(id, name, userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllListingsResponse{
		Data: listings,
	})
}

func (h *Handler) createListing(ctx *gin.Context) {
	var listing model.Listing

	if err := ctx.ShouldBindJSON(&listing); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	listingId, err := h.services.Listing.CreateListing(listing)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": listingId,
	})
}

func (h *Handler) updateListing(ctx *gin.Context) {
	var listing model.Listing

	if err := ctx.ShouldBindJSON(&listing); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := listing.ID
	name := listing.Name
	fmt.Print("id = ", id, "\n")
	fmt.Print("name = ", name, "\n")

	userId, err := h.services.Listing.UpdateListing(id, name)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"user_id": userId,
	})
}

func (h *Handler) deleteListing(ctx *gin.Context) {
	var listing model.Listing

	if err := ctx.ShouldBindJSON(&listing); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := listing.ID

	listingId, err := h.services.Listing.DeleteListing(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": listingId,
	})
}
