package handler

import (
	"net/http"
	"web/internal/model"

	"github.com/gin-gonic/gin"
)

type getAllListingsResponse struct {
	Data []model.Listing `json:"data"`
}

func (h *Handler) getListingsByName(ctx *gin.Context) {
	listings, err := h.services.Listing.GetListingsByName()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllListingsResponse{
		Data: listings,
	})
}

// func (h *Handler) createListing(ctx *gin.Context) {
// }
