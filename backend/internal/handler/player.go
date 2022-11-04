package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPlayer(ctx *gin.Context) {
	// id, ok := ctx.Get(usertCtx)

	// if !ok {
	// 	newErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
	// 	return
	// }

	// var input model.Player
	// if err := ctx.BindJSON(&input); err != nil {
	// 	newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// call service method
}

func (h *Handler) getAllPlayers(ctx *gin.Context) {

}

func (h *Handler) getPlayerById(ctx *gin.Context) {

}

func (h *Handler) updatePlayer(ctx *gin.Context) {

}

func (h *Handler) deletePlayer(ctx *gin.Context) {

}
