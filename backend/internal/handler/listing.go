package handler

import (
	"web/internal/model"

	"github.com/gin-gonic/gin"
)

type getAllListingsResponse struct {
	Data []model.Listing `json:"data"`
}

func (h *Handler) getAllListings(ctx *gin.Context) {

}
