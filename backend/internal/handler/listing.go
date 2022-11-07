package handler

import (
	"fmt"
	"net/http"
	"web/internal/model"

	"github.com/gin-gonic/gin"
)

type getAllListingsResponse struct {
	Data []model.Listing `json:"data"`
}

func (h *Handler) getListingsByName(ctx *gin.Context) {
	name := ctx.Query("name")
	// if err != nil {
	// 	newErrorResponse(ctx, http.StatusBadRequest, "invalid name param")
	// 	return
	// }

	fmt.Print("name = ", name, "\n")

	listings, err := h.services.Listing.GetListingsByName(name)
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
