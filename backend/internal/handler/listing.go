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

	// if err := ctx.ShouldBindJSON(&json); err != nil {
	// 	newErrorResponse(ctx, http.StatusBadRequest, "invalid json")
	// 	return
	// }

	if err := ctx.ShouldBindWith(&listing, binding.Query); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := listing.ID
	name := listing.Name
	user_id := listing.UserID

	fmt.Print("id = ", id, "\n")
	fmt.Print("name = ", name, "\n")
	fmt.Print("user_id = ", user_id, "\n")

	listings, err := h.services.Listing.GetListings(id, name, user_id)
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
